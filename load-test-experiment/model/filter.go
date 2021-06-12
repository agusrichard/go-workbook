package model

type Query struct {
	Skip    int      `form:"skip" json:"skip" xml:"skip"`
	Take    int      `form:"take" json:"take" xml:"take"`
	Orders  []Order  `form:"orders" json:"orders" xml:"orders"`
	Search  string   `form:"search" json:"search" xml:"search"`
	Filters []Filter `form:"filters" json:"filters" xml:"filters"`
}

type Filter struct {
	Type  string `form:"type" json:"type" xml:"type"`
	Value string `form:"value" json:"value" xml:"value"`
	Field string `form:"field" json:"field" xml:"field"`
	Start string `form:"start" json:"start" xml:"start"`
	End   string `form:"end" json:"end" xml:"end"`
}

type Order struct {
	Field string // `form:"field" json:"field" xml:"field"`
	Dir   string // `form:"dir" json:"dir" xml:"dir"`
}
