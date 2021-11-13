package service

import (
	"github.com/glavanan/talanddev/dbconnector"
	"github.com/glavanan/talanddev/model"
)

//AddHistory save a convertion in DB
func AddHistory(history model.History) error {
	db, err := dbconnector.GetDB()
	if err != nil {
		return err
	}
	err = db.Create(&history).Error
	return err
}

//GetHistories Get histories saved in DB
func GetHistories() ([]model.History, error) {
	db, err := dbconnector.GetDB()
	if err != nil {
		return nil, err
	}
	var histories []model.History
	err = db.Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}
