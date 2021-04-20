package repo

const sqlGetReviewsForDate = `
SELECT reviewer, repo, pr_id, badge from reviews
WHERE date(ts) = date(?)
`
const sqlGetReviewsForWeek = `
SELECT reviewer, repo, pr_id, badge from reviews
WHERE WEEKOFYEAR(ts) = WEEKOFYEAR(?) AND year(ts) = year(?)
`

const sqlNewReview = `
INSERT INTO reviews (reviewer,repo, pr_id, badge) 
VALUES (?,?,?,?)
`
const sqlPing = `
SELECT "Database is online"
`
