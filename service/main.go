package service

import (
	"errors"
	"fmt"
	"github.com/celerway/diamonds/dbos"
	"github.com/celerway/diamonds/dtos"
	"github.com/celerway/diamonds/repo"
	log "github.com/sirupsen/logrus"
	"time"
)

type DiamondService struct {
	Repo repo.Repository
}

func Initialize(repo repo.Repository) DiamondService {
	s := DiamondService{Repo: repo}
	log.Info("Service layer initialized")
	return s
}

func (s DiamondService) Ping() (string, error) {
	return s.Repo.Ping()
}

func (s DiamondService) RegisterReview(review dtos.Review) error {

	reviewDb := dbos.Review{
		Reviewer: review.Reviewer,
		Repo:     review.Repo,
		Pr:       review.Pr,
		Badge:    review.Badge,
	}
	err := s.Repo.RegisterReview(reviewDb)
	if err != nil {
		return err
	}
	return nil
}

func (s DiamondService) GetStats(period string, when time.Time) (dtos.ReviewStatMap, error) {
	reviews := make(dtos.ReviewStatMap, 0) // the response map.
	var dbReviews dbos.Reviews
	var err error

	switch period {
	case "week":
		dbReviews, err = s.Repo.GetStatsForWeek(when)
	case "day":
		dbReviews, err = s.Repo.GetStatsForDay(when)
	default:
		return nil, errors.New(fmt.Sprintf("Unknown period specification: %s", period))
	}

	if err != nil {
		return nil, err
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

	return reviews, err

}
