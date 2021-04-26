package api

import (
	"github.com/gorilla/mux"
	"net/http"
)


const baseUrl = "/diamonds/v1/"

func makeRoutes(app App) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(baseUrl, app.HomeHandler).Methods("GET")
	r.HandleFunc(baseUrl + "review", app.ReviewHandler).Methods("POST")
	r.HandleFunc(baseUrl + "reviews", app.ReviewsHandler).Methods("GET")
	r.HandleFunc("/healthz", app.HomeHandler).Methods("GET") // Can work as a healthz handler as well. Pings the db.
	http.Handle("/", r)
	return r
}
