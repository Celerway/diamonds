package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func makeRoutes(app App) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", app.HomeHandler).Methods("GET")
	r.HandleFunc("/v1/review", app.ReviewHandler).Methods("POST")
	r.HandleFunc("/v1/reviews", app.ReviewsHandler).Methods("GET")
	http.Handle("/", r)
	return r
}
