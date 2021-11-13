package model

import "time"

//Convert request to convert an amount with another currency
type Convert struct {
	Amount    float64   `json:"amount"`
	Date      time.Time `json:"date"`
	Currency  string    `json:"currency"`
	ConvertTo string    `json:"convert_to"`
}

//Exchange map of value of currency
type Exchange struct {
	Rates map[string]float64 `json:"rates"`
}

//History model of data in db
type History struct {
	ID           int       `json:"id" gorm:"id"`
	Amount       float64   `json:"amount" gorm:"amount"`
	CurrencyFrom string    `json:"currency_from" gorm:"currency_from"`
	CurrencyTo   string    `json:"currency_to" gorm:"currency_to"`
	Result       float64   `json:"result" gorm:"result"`
	Date         time.Time `json:"date" gorm:"date"`
}

//TableName ...
func (h History) TableName() string {
	return "history"
}
