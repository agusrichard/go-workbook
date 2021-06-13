package model

import (
	"database/sql"
	"time"
)

type MediumV1SmallModel struct {
	ID            int            `json:"id" db:"id"`
	FieldOne      string         `json:"fieldOne" db:"field_one"`
	FieldTwo      float64        `json:"fieldTwo" db:"field_two"`
	FieldThree    sql.NullString `json:"fieldThree" db:"field_three"`
	FieldFour     sql.NullTime   `json:"fieldFour" db:"field_four"`
	SmallLargeKey int            `json:"smallLargeKey" db:"small_large_key"`
}

type MediumV1Model struct {
	ID                   int                  `json:"id" db:"id"`
	FieldOne             string               `json:"fieldOne" db:"field_one"`
	FieldTwo             float64              `json:"fieldTwo" db:"field_two"`
	FieldThree           sql.NullString       `json:"fieldThree" db:"field_three"`
	FieldFour            time.Time            `json:"fieldFour" db:"field_four"`
	FieldFive            int                  `json:"FieldFive" db:"field_five"`
	FieldSix             string               `json:"fieldSix" db:"field_six"`
	FieldSeven           float64              `json:"fieldSeven" db:"field_seven"`
	FieldEight           sql.NullString       `json:"fieldEight" db:"field_eight"`
	FieldNine            time.Time            `json:"fieldNine" db:"field_nine"`
	FieldTen             int                  `json:"fieldTen" db:"field_ten"`
	FieldEleven          string               `json:"fieldEleven" db:"field_eleven"`
	FieldTwelve          float64              `json:"fieldTwelve" db:"field_twelve"`
	FieldThirteen        sql.NullString       `json:"fieldThirteen" db:"field_thirteen"`
	FieldFourteen        time.Time            `json:"fieldFourteen" db:"field_fourteen"`
	MediumSmallModelList []MediumV1SmallModel `json:"mediumSmallModelList"`
}

type MediumV1SmallShape struct {
	ID            int       `json:"id" db:"id"`
	FieldOne      string    `json:"fieldOne" db:"field_one"`
	FieldTwo      float64   `json:"fieldTwo" db:"field_two"`
	FieldThree    string    `json:"fieldThree" db:"field_three"`
	FieldFour     string `json:"fieldFour" db:"field_four"`
	SmallLargeKey int       `json:"smallLargeKey" db:"small_large_key"`
}

type MediumV1Shape struct {
	ID                   int                  `json:"id" db:"id"`
	FieldOne             string               `json:"fieldOne" db:"field_one"`
	FieldTwo             float64              `json:"fieldTwo" db:"field_two"`
	FieldThree           string           `json:"fieldThree" db:"field_three"`
	FieldFour            time.Time            `json:"fieldFour" db:"field_four"`
	FieldFive            int                  `json:"FieldFive" db:"field_five"`
	FieldSix             string               `json:"fieldSix" db:"field_six"`
	FieldSeven           float64              `json:"fieldSeven" db:"field_seven"`
	FieldEight           string           `json:"fieldEight" db:"field_eight"`
	FieldNine            time.Time            `json:"fieldNine" db:"field_nine"`
	FieldTen             int                  `json:"fieldTen" db:"field_ten"`
	FieldEleven          string               `json:"fieldEleven" db:"field_eleven"`
	FieldTwelve          float64              `json:"fieldTwelve" db:"field_twelve"`
	FieldThirteen        string           `json:"fieldThirteen" db:"field_thirteen"`
	FieldFourteen        time.Time            `json:"fieldFourteen" db:"field_fourteen"`
	MediumSmallModelList []MediumV1SmallShape `json:"mediumSmallModelList"`
}

type MediumV1Request struct {
	Request
	Data MediumV1Shape
}

type MediumV1Response struct {
	Data    []MediumV1Shape `json:"data"`
	Message string          `json:"message"`
}
