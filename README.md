# 💎 Diamonds 💎

This is a somewhat silly tool to hand out emojis to people participating in software development. The idea is that you
create Github triggers that would call the /v1/review endpoint and register that someone just did something awesome.

At certain points in time the daemon connects to slack and posts stats about who got what that day.

At the end of the week it will also produce a weekly summary.

## Required environment variables

The daemon needs the following environment values to be set:

- DSN - database connection string. MySQL is the only supported database other database are likely just a recompile
  away.
- APIKEY - The API key required to post to the HTTP interface.
- SLACK_CHANNELID - The channel where you want the bot to report
- SLACK_TOKEN - You slack bot token.

## Usage

In order to give "bob" a 🦆 for his wonderful work on reviewing duck typing you in PR 5 can run the following command (
httpie).

```
 http -v  'localhost:4210/v1/review' APIKey:mysecretkey Reviewer=bob Repo=github.com/celerway/diamonds 'Pr:=5' Badge=🦆
```

This will create the following POST body:

```
{
    "Badge": "🦆",
    "Pr": 15,
    "Repo": "github.com/celerway/diamonds",
    "Reviewer": "bob"
}
```

## Libraries used:

- Gorilla Mux
- SQLx
- logrus – See go.mod for a complete list.

## Todo

- There are no tests as of now.
- No helm
- Listens to :4210 (hard coded)  
- Require a Content-type header with application/json
- The scheduler is hardcoded with consts.  
- Probably a lot more

PRs are welcome. 