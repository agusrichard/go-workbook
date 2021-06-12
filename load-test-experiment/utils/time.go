package util

import (
	"fmt"
	"time"
)

func ParseTime(strTime string) (time.Time, error) {
	if strTime == "" {
		return time.Time{}, nil
	}

	t, err := time.Parse(time.RFC3339, strTime)
	if err != nil {
		fmt.Println("error parse time", err)
		return time.Time{}, err
	}
	return t, nil
}

func TimeToStr(t time.Time) (string, error) {
	return "", nil
}