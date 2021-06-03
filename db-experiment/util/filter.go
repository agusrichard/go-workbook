package util

import (
	model "db-experiment/models"
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"time"
)

func CreateQueryFilter(filters *[]model.Filter, additionals *[]model.Filter) (string, error) {
	var toBeIterated []model.Filter
	var clauseList []string

	if filters == nil {
		return "", errors.New("create query filter: filter should not be nil")
	}

	if additionals == nil {
		toBeIterated = *filters
	} else {
		toBeIterated = append(*filters, *additionals...)
	}

	for _, v := range toBeIterated {
		var clause string

		switch v.Type {
		case "text":
			clause = "lower(" + v.Field + `) like lower('%` + v.Value + "%') "
		case "date":
			fdMin, err := FilterDate(v.Value, "min")
			if err != nil {
				return "", errors.New("create query filter: error create filter date min")
			}

			fdMax, err := FilterDate(v.Value, "max")
			if err != nil {
				return "", errors.New("create query filter: error create filter date max")
			}

			clause = v.Field + `>='` + fdMin + "'" + " AND " + v.Field + `<'` + fdMax + "'"
		default:
			clause = v.Field + `=` + v.Value
		}

		clauseList = append(clauseList, clause)
	}

	filterString := strings.Join(clauseList, " AND ")

	return filterString, nil
}

func FilterDate(date, types string) (string, error) {
	var timeAdd time.Time

	timeStr := fmt.Sprintf(`%sT00:00:00.000Z`, date)
	timeParse, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return "", errors.New("filter date: error parse string to time")
	}

	if types == "min" {
		timeAdd = timeParse.Add(-7 * time.Hour)
	} else {
		timeAdd = timeParse.Add(17 * time.Hour)
	}

	return TimeToString(timeAdd), nil
}

func TimeToString(date time.Time) string {
	dateString := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.000Z",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second())

	return dateString
}