package dbos

type Review struct {
	// Id       int    `db:"id"`  // We never specify when we write.
	Reviewer string `db:"reviewer"`
	Repo     string `db:"repo"`
	Pr       int    `db:"pr_id"`
	Badge    string `db:"badge"`
}

type Reviews []Review
