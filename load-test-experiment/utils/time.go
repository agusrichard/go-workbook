package util

import (
	"time"
)

func ParseTime(strTime string) (time.Time, error) {
	if strTime == "" {
		return time.Time{}, nil
	}

	t, err := time.Parse(time.RFC3339, strTime)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func TimeToStr(t time.Time) (string, error) {
	return "", nil
}