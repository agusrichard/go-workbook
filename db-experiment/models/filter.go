package model

type Filter struct {
	Type  string `form:"type" json:"type" xml:"type"`
	Value string `form:"value" json:"value" xml:"value"`
	Field string `form:"field" json:"field" xml:"field"`
	Start string `form:"start" json:"start" xml:"start"`
	End   string `form:"end" json:"end" xml:"end"`
}

type Query struct {
	Skip         int    `form:"skip" json:"skip" xml:"skip"`
	Take         int    `form:"take" json:"take" xml:"take"`
	Order        string `form:"order" json:"order" xml:"order"`
	Search       string `form:"search" json:"search" xml:"search"`
	FilterString string `form:"filter" json:"filter" xml:"filter"`
}
