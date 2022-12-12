package main

import (
	"testing"
	"time"
)

func TestTimeInLog(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected time.Time
	}

	tests := []test{
		{
			name:     "one",
			input:    `127.0.0.1 user-identifier frank [06/Dec/2021:18:49:17 +0000] "GET /api/endpoint HTTP/1.0" 200 5134`,
			expected: time.Date(2021, time.December, 6, 18, 49, 17, 0, time.UTC),
		},
		{
			name:     "two",
			input:    `127.0.0.1 user-identifier frank [13/Dec/2021:09:00:00 +0000] "GET /api/endpoint HTTP/1.0" 200 5134`,
			expected: time.Date(2021, time.December, 13, 9, 0, 0, 0, time.UTC),
		},
		{
			name:     "three",
			input:    `127.0.0.1 user-identifier frank [21/Sep/1997:07:00:00 +0000] "GET /api/endpoint HTTP/1.0" 200 5134`,
			expected: time.Date(1997, time.September, 21, 7, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := findTimeInLog(tc.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkAnalyze(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minutes := "10m"
		dir := "./logs"
		analyze(&minutes, &dir)
	}
}
