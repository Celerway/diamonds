# Diamonds!


This is a somewhat silly tool to hand out emojis to people participating in 
software development. The idea is that you create Github triggers that 
would call the /v1/review endpoint and register that someone just did
something awesome.

At certain points in time the daemon connects to slack and posts stats
about who got what that day.

At the end of the week it will also produce a weekly summary.

## Usage
The daemon needs the following environment values to be set:
DSN - database connection string. MySQL is the only supported database other database are likely just a recompile away.
APIKEY - The API key required.
SLACK_URL - Your SLACK URL endpoint. So we can talk to Slack. 

## Libraries used:
 - Gorilla Mux
 - SQLx
 - logrus
 – See go.mod for a complete list.

## Todo
 - There are no tests as of now.
