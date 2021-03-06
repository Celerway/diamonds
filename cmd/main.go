package main

import (
	"github.com/celerway/diamonds/api"
	"github.com/celerway/diamonds/repo"
	"github.com/celerway/diamonds/scheduler"
	"github.com/celerway/diamonds/service"
	"github.com/celerway/diamonds/slapp"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)
	err := godotenv.Load()
	if err != nil {
		log.Infof("Error loading .env file, assuming production: %s", err.Error())
	}
	myRepo := repo.Initialize()
	myService := service.Initialize(myRepo)
	mySlapp := slapp.Initialize(myService)
	myApi := api.Initialize(myService)
	// glue the slack app to the scheduler:
	scheduler.Initialize(myService, mySlapp)
	myApi.Run() // starts the API. Blocks.
}
