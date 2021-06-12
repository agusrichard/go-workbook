package model

import (
	"database/sql"
)

type LightV1Model struct {
	ID         int            `json:"id"`
	FieldOne   string         `json:"fieldOne"`
	FieldTwo   float64        `json:"fieldTwo"`
	FieldThree sql.NullString `json:"fieldThree"`
	FieldFour  sql.NullTime   `json:"fieldFour"`
}

type LightV1Shape struct {
	ID         int     `json:"id"`
	FieldOne   string  `json:"fieldOne"`
	FieldTwo   float64 `json:"fieldTwo"`
	FieldThree string  `json:"fieldThree"`
	FieldFour  string  `json:"fieldFour" time_format:"2006-01-02T15:04:05Z07:00"`
}

type LightV1Request struct {
	Request
	Data LightV1Shape
}
