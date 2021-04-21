package repo

import (
	"github.com/celerway/diamonds/dbos"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type Repository struct {
	db *sqlx.DB
}

func Initialize() Repository {
	r := Repository{}
	db, err := sqlx.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}
	r.db = db
	log.Info("Repo layer initialized")
	return r
}

func (repo Repository) Ping() (string, error) {
	var res string
	err := repo.db.Get(&res, sqlPing)
	return res, err
}

func (repo Repository) RegisterReview(r dbos.Review) error {
	_, err := repo.db.Exec(sqlNewReview,
		r.Reviewer, r.Repo, r.Pr, r.Badge)
	return err
}

func (repo Repository) GetStatsForDay(when time.Time) (dbos.Reviews, error) {
	var reviews dbos.Reviews
	err := repo.db.Select(&reviews, sqlGetReviewsForDate, when)
	if err != nil {
		return nil, err
	}
	return reviews, err
}

func (repo Repository) GetStatsForWeek(when time.Time) (dbos.Reviews, error) {
	var reviews dbos.Reviews
	err := repo.db.Select(&reviews, sqlGetReviewsForWeek, when, when)
	if err != nil {
		return nil, err
	}
	return reviews, err
}
