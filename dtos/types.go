package dtos

// Review
// For registering a review or viewing a single review
type Review struct {
	Reviewer string `json:"Reviewer"`
	Repo     string `json:"Repo"`
	Pr       int    `json:"Pr"`
	Badge    string `json:"Badge"`
}

type PrMap map[int]string // Map with pr --> URL.

type ReviewerStat struct {
	// Reviewer string // Not needed. It is the key.
	Prs    PrMap
	Badges string
}

// ReviewStatMap
// The response to a query for a reviews for a day or week.
type ReviewStatMap map[string]ReviewerStat
