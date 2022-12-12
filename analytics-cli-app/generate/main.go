package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"time"
)

const (
	NUM_OF_FILES = 10
	NUM_OF_LINES = 100000
	LOG_FORMAT   = "127.0.0.1 user-identifier frank [%s] \"GET /api/endpoint HTTP/1.0\" 200 5134\n"
)

func main() {
	generate(NUM_OF_FILES, NUM_OF_LINES)
}

func generate(numOfFiles, numOfLines int) {
	fmt.Println("Generating log data...")

	logsDir := getLogsDir()

	_, err := os.Stat(logsDir)
	if !os.IsNotExist(err) {
		clearLogDir()
	}

	err = os.Mkdir(logsDir, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	currentTime := time.Now()

	for i := numOfFiles; i >= 1; i-- {
		f, err := os.Create(path.Join(logsDir, fmt.Sprintf("http-%d.log", i)))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lines := createLogLines(currentTime, numOfLines)

		for _, line := range lines {
			_, err := f.WriteString(line)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		time.Sleep(time.Millisecond * 10)
		f.Close()
	}
	fmt.Println("Log data has been generated")
}

func clearLogDir() {
	fmt.Println("Clearing log directory...")
	logsDir := getLogsDir()
	err := os.RemoveAll(logsDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Log directory has been cleared")
}

func getLogsDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootDir := filepath.Dir(pwd)

	return path.Join(rootDir, "analytics", "logs")
}

func createLogLines(currentTime time.Time, numOfLines int) []string {
	lines := make([]string, 0)
	for j := 0; j < numOfLines; j++ {
		timeFormatted := currentTime.Format("02/Jan/2006:15:04:05 +0000")

		lines = append(lines, fmt.Sprintf(LOG_FORMAT, timeFormatted))

		currentTime = currentTime.Add(time.Minute * -1)
	}

	reverse(lines)

	return lines
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
