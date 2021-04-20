package api

import (
	"github.com/celerway/diamonds/repo"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	Repo   *repo.Repository
}

func Initialize() App {
	app := App{
		Repo: repo.Initialize(),
	}
	app.Router = makeRoutes(app)
	app.Router.Use(AuthMiddleware)
	return app
}

func (app App) Run() {
	log.Fatal(http.ListenAndServe(":4210", app.Router))
}
