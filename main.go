package main

import (
	"github.com/glavanan/talanddev/dbconnector"
	"github.com/glavanan/talanddev/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	//The authentication should have been in a config file.
	dbconnector.InitDB("host=0.0.0.0 user=postgres password=secretpass dbname=talanddev port=5432")
	router := router.GetRouter()
	err := router.Run(":" + "9090")
	if err != nil {
		log.Errorf("Error when running the router")
	}
}
