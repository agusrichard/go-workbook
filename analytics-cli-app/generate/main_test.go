package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected []int
	}

	tests := []test{
		{
			name:     "one to five",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "one to ten",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:     "one to twenty",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			expected: []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := make([]int, len(tc.input))
			copy(actual, tc.input)
			reverse(actual)
			for i := 0; i < len(tc.expected); i++ {
				if actual[i] != tc.expected[i] {
					t.Errorf("Expected %d, got %d", tc.expected[i], actual[i])
				}
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	type test struct {
		name     string
		input    []string
		expected []string
	}

	tests := []test{
		{
			name:     "one to five",
			input:    []string{"one", "two", "three", "four", "five"},
			expected: []string{"five", "four", "three", "two", "one"},
		},
		{
			name:     "one to ten",
			input:    []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
			expected: []string{"ten", "nine", "eight", "seven", "six", "five", "four", "three", "two", "one"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := make([]string, len(tc.input))
			copy(actual, tc.input)
			reverse(actual)
			for i := 0; i < len(tc.expected); i++ {
				if actual[i] != tc.expected[i] {
					t.Errorf("Expected %s, got %s", tc.expected[i], actual[i])
				}
			}
		})
	}
}

func setupTest(t *testing.T) func(t *testing.T) {
	log.Println("Setup test")

	return func(t *testing.T) {
		log.Println("Teardown test")
	}
}

func TestGenerate(t *testing.T) {
	defer clearLogDir()

	type nums struct {
		numOfFiles int
		numOfLines int
	}

	tests := []struct {
		name  string
		input nums
	}{
		{
			name: "1 file 1 line",
			input: nums{
				numOfFiles: 1,
				numOfLines: 1,
			},
		},
		{
			name: "5 files 5 lines",
			input: nums{
				numOfFiles: 5,
				numOfLines: 5,
			},
		},
		{
			name: "10 files 10 lines",
			input: nums{
				numOfFiles: 10,
				numOfLines: 10,
			},
		},
		{
			name: "100 files 100 lines",
			input: nums{
				numOfFiles: 100,
				numOfLines: 100,
			},
		},
		{
			name: "1000 files 1000 lines",
			input: nums{
				numOfFiles: 1000,
				numOfLines: 1000,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			teardown := setupTest(t)
			defer teardown(t)
			generate(tc.input.numOfFiles, tc.input.numOfLines)

			logsDir := getLogsDir()
			files, err := ioutil.ReadDir(logsDir)
			if err != nil {
				t.Errorf("Failed to read logs directory: %s", err)
			}

			if len(files) != tc.input.numOfFiles {
				t.Errorf("Expected %d files, got %d", tc.input.numOfFiles, len(files))
			}

			for _, file := range files {
				filePath := path.Join(logsDir, file.Name())
				file, _ := os.Open(filePath)
				fileScanner := bufio.NewScanner(file)
				lineCount := 0
				for fileScanner.Scan() {
					lineCount++
				}
				// fmt.Println("number of lines:", lineCount)
				if lineCount != tc.input.numOfLines {
					t.Errorf("Expected %d lines, got %d", tc.input.numOfLines, lineCount)
				}
			}
		})
	}
}

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(100, 100)
	}

	clearLogDir()
}
