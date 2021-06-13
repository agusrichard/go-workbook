package model

type MediumV2SmallModel struct {
	ID         int        `json:"id" db:"id"`
	FieldOne   string     `json:"fieldOne" db:"field_one"`
	FieldTwo   float64    `json:"fieldTwo" db:"field_two"`
	FieldThree NullString `json:"fieldThree" db:"field_three"`
	FieldFour  NullTime   `json:"fieldFour" db:"field_four"`
}

type MediumV2LargeModel struct {
	ID            int        `json:"id" db:"id"`
	FieldOne      string     `json:"fieldOne" db:"field_one"`
	FieldTwo      float64    `json:"fieldTwo" db:"field_two"`
	FieldThree    NullString `json:"fieldThree" db:"field_three"`
	FieldFour     NullTime   `json:"fieldFour" db:"field_four"`
	FieldFive     int        `json:"FieldFive" db:"field_five"`
	FieldSix      string     `json:"fieldSix" db:"field_six"`
	FieldSeven    float64    `json:"fieldSeven" db:"field_seven"`
	FieldEight    NullString `json:"fieldEight" db:"field_eight"`
	FieldNine     NullTime   `json:"fieldNine" db:"field_nine"`
	FieldTen      int        `json:"fieldTen" db:"field_ten"`
	FieldEleven   string     `json:"fieldEleven" db:"field_eleven"`
	FieldTwelve   float64    `json:"fieldTwelve" db:"field_twelve"`
	FieldThirteen NullString `json:"fieldThirteen" db:"field_thirteen"`
	FieldFourth   NullTime   `json:"fieldFourth" db:"field_fourth"`
	SmallLargeKey int        `json:"smallLargeKey" db:"small_large_key"`
}
