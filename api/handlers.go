package api

import (
	"encoding/json"
	"fmt"
	"github.com/celerway/diamonds/dbos"
	"github.com/celerway/diamonds/dtos"
	"io"
	"net/http"
	"time"
)

type Hello struct {
	Message string `json:"Message"`
}

func (app App) HomeHandler(w http.ResponseWriter, _ *http.Request) {

	resp, err := app.Repo.Ping()
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)

	reviewDb := dbos.Review{
		Reviewer: review.Reviewer,
		Repo:     review.Repo,
		Pr:       review.Pr,
		Badge:    review.Badge,
	}

	err := app.Repo.RegisterReview(reviewDb)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, review)
}

// ReviewsHandler
// Responds with the reviews given
func (app App) ReviewsHandler(w http.ResponseWriter, r *http.Request) {
	reviews := make(dtos.ReviewStatMap, 0) // the response map.
	var dbReviews dbos.Reviews
	var err error

	v := r.URL.Query()
	period := v.Get("period")

	if period == "" {
		period = "day"
	}

	switch period {
	case "week":
		dbReviews, err = app.Repo.GetStatsForWeek(time.Now().UTC())
	case "day":
		dbReviews, err = app.Repo.GetStatsForDay(time.Now().UTC())
	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Unknown period specification: %s", period))
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	for _, review := range dbReviews {
		// Get the current reviewer from the map.
		curReviewer, ok := reviews[review.Reviewer]
		if !ok { // if it doesn't exist, make one.
			curReviewer = dtos.ReviewerStat{}
			curReviewer.Prs = make([]string, 0)
		}
		// Add the badge
		curBadges := curReviewer.Badges + review.Badge
		curReviewer.Badges = curBadges
		// Add the PR
		curReviewer.Prs = append(curReviewer.Prs, fmt.Sprintf("%s/pulls/%d", review.Repo, review.Pr))
		// Write it back to the map.
		reviews[review.Reviewer] = curReviewer
	}
	respondWithJSON(w, http.StatusOK, reviews)
}
