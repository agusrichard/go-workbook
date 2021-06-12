package model

type Request struct {
	Action string `json:"action"`
	Query  Query  `json:"query"`
}
