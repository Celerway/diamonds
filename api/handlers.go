package api

import (
	"encoding/json"
	"fmt"
	"github.com/celerway/diamonds/dtos"
	"io"
	"net/http"
	"time"
)

type Hello struct {
	Message string `json:"Message"`
}

func (app App) HomeHandler(w http.ResponseWriter, _ *http.Request) {

	resp, err := app.Service.Ping()
	h := Hello{resp}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, h)
	}
}

// ReviewHandler
// Accepts a new review.
func (app App) ReviewHandler(w http.ResponseWriter, r *http.Request) {
	var review dtos.Review

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&review); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %s", err))
		return
	}

	// The somewhat ugly way of handling an error in a defer statement:
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)

	err := app.Service.RegisterReview(review)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, review)
}

// ReviewsHandler
// Responds with the reviews given
func (app App) ReviewsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	reviews := make(dtos.ReviewStatMap, 0) // the response map.

	v := r.URL.Query()
	period := v.Get("period")

	if period == "" {
		period = "day"
	}

	reviews, err = app.Service.GetStats(period, time.Now())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, reviews)
}
