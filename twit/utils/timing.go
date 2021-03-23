package utils

import (
	"log"
	"time"
)

func MeasureExecutionTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
