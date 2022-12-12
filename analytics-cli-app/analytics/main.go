package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	REGEX_PATTERN = `\d{1,2}\/(Jan(uary)?|Feb(ruary)?|Mar(ch)?|Apr(il)?|May|Jun(e)?|Jul(y)?|Aug(ust)?|Sep(tember)?|Oct(ober)?|Nov(ember)?|Dec(ember)?)\/\d{4}:\d{1,2}:\d{1,2}:\d{1,2}\s+\+\d{4}`
)

var ErrTMinutesReached = errors.New("ErrTMinutesReached")

func main() {
	fmt.Printf("========== Processing log data ========== \n\n")

	// Specifying command options to run the cli application
	tMinutesFlag := flag.String("t", "", "time in minutes")
	directoryFlag := flag.String("d", "", "directory")
	flag.Parse()

	// Make sure that the required flags are set
	if *tMinutesFlag == "" && *directoryFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	analyze(tMinutesFlag, directoryFlag)
	fmt.Printf("\n========== End of processing ==========")
}

func analyze(tMinutesFlag, directoryFlag *string) {

	// Parse the time in minutes
	tMinutes, err := strconv.ParseInt(strings.Replace(*tMinutesFlag, "m", "", 1), 10, 64)
	if err != nil {
		os.Exit(1)
	}

	var lastRecordedTime time.Time
	firstRecord := true

	dirPath := *directoryFlag
	files, err := readDir(dirPath)
	for _, file := range files {
		err := processFile(file, tMinutes, &lastRecordedTime, &firstRecord)
		if err == ErrTMinutesReached {
			fmt.Printf("\n========== End of processing ==========\n")
			return
		}
	}
}

func readDir(dirpath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Before(files[j].ModTime())
	})

	result := make([]string, 0)
	for _, file := range files {
		result = append(result, path.Join((dirpath), file.Name()))
	}

	return result, nil
}

func processFile(filepath string, lastTMinutes int64, lastRecordedTime *time.Time, firstRecord *bool) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	f1, err := f.Stat()
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := NewScanner(f, int(f1.Size()))
	for {
		line, _, err := scanner.Line()
		if err != nil {
			return err
		}

		if line == "" {
			continue
		}

		if *firstRecord {
			t, err := findTimeInLog(line)
			if err != nil {
				return err
			}

			*lastRecordedTime = t.Add(time.Minute * time.Duration(lastTMinutes) * -1)
			*firstRecord = false
		}

		t, err := findTimeInLog(line)
		if t.Before(*lastRecordedTime) {
			return ErrTMinutesReached
		}

		fmt.Println(line)
	}
}

func findTimeInLog(logstr string) (time.Time, error) {
	re := regexp.MustCompile(REGEX_PATTERN)

	timeStr := re.FindString(logstr)
	t, err := time.Parse("02/Jan/2006:15:04:05 +0000", timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

type Scanner struct {
	r   io.ReaderAt
	pos int
	err error
	buf []byte
}

func NewScanner(r io.ReaderAt, pos int) *Scanner {
	return &Scanner{r: r, pos: pos}
}

func (s *Scanner) readMore() {
	if s.pos == 0 {
		s.err = io.EOF
		return
	}
	size := 1024
	if size > s.pos {
		size = s.pos
	}
	s.pos -= size
	buf2 := make([]byte, size, size+len(s.buf))

	// ReadAt attempts to read full buff!
	_, s.err = s.r.ReadAt(buf2, int64(s.pos))
	if s.err == nil {
		s.buf = append(buf2, s.buf...)
	}
}

func (s *Scanner) Line() (line string, start int, err error) {
	if s.err != nil {
		return "", 0, s.err
	}
	for {
		lineStart := bytes.LastIndexByte(s.buf, '\n')
		if lineStart >= 0 {
			// We have a complete line:
			var line string
			line, s.buf = string(dropCR(s.buf[lineStart+1:])), s.buf[:lineStart]
			return line, s.pos + lineStart + 1, nil
		}
		// Need more data:
		s.readMore()
		if s.err != nil {
			if s.err == io.EOF {
				if len(s.buf) > 0 {
					return string(dropCR(s.buf)), 0, nil
				}
			}
			return "", 0, s.err
		}
	}
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
