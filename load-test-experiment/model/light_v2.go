package model

type LightV2Model struct {
	ID         int        `json:"id" db:"id"`
	FieldOne   string     `json:"fieldOne" db:"field_one"`
	FieldTwo   float64    `json:"fieldTwo" db:"field_two"`
	FieldThree NullString `json:"fieldThree" db:"field_three"`
	FieldFour  NullTime   `json:"fieldFour" db:"field_four"`
}

type LightV2Request struct {
	Request
	Data LightV2Model
}
