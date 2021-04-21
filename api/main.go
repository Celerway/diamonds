package api

import (
	"github.com/celerway/diamonds/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	log.Fatal(http.ListenAndServe(":4210", app.Router))
}
