package main

import (
	"github.com/celerway/diamonds/api"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Infof("Error loading .env file, assuming production: %s", err.Error())
	}
	app := api.Initialize()
	app.Run()
}
