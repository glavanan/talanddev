package model

import "time"

//Convert request to convert an amount with another currency
type Convert struct {
	Amount    float64   `json:"amount"`
	Date      time.Time `json:"date"`
	Currency  string    `json:"currency"`
	ConvertTo string    `json:"convert_to"`
}

type Exchange struct {
	Rates map[string]float64 `json:"rates"`
}
