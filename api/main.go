package api

import (
	"fmt"
	"github.com/celerway/diamonds/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type App struct {
	Router  *mux.Router
	Service service.DiamondService
}

func Initialize(service service.DiamondService) App {
	app := App{
		Service: service,
	}
	app.Router = makeRoutes(app)
	app.Router.Use(AuthMiddleware)
	log.Info("API initialized")
	return app
}

func (app App) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4210"
	}
	port = fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(port, app.Router))
}
