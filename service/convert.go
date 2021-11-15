package service

import (
	"errors"

	"github.com/glavanan/talanddev/dbconnector"
	"github.com/glavanan/talanddev/model"
)

//Convert the amount with the given currency
func Convert(conv model.Convert) (*model.History, error) {
	exchange := dbconnector.GetExchangeRate()
	divide, ok := exchange[conv.Currency]
	if !ok || divide == 0 {
		return nil, errors.New("Couldn't find the current currency")
	}
	toConv := conv.Amount / divide
	multiply, ok := exchange[conv.ConvertTo]
	if !ok || multiply == 0 {
		return nil, errors.New("Couldn't find the currency to conver to")
	}
	result := model.History{Date: conv.Date,
		CurrencyFrom: conv.Currency,
		CurrencyTo:   conv.ConvertTo,
		Amount:       conv.Amount,
		Result:       toConv * multiply,
	}
	err := AddHistory(result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
