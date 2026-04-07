package model

type Transaction struct {
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
	Desc   string  `json:"desc"`
}
