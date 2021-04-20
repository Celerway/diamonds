package dtos

// Review
// For registering a review or viewing a single review
type Review struct {
	Reviewer string `json:"Reviewer"`
	Repo     string `json:"Repo"`
	Pr       int    `json:"Pr"`
	Badge    string `json:"Badge"`
}

type ReviewerStat struct {
	// Reviewer string // Not needed. It is the key.
	Prs    []string // Contains URLs to all the Prs
	Badges string
}

// ReviewStatMap
// The response to a query for a reviews for a day or week.
type ReviewStatMap map[string]ReviewerStat
