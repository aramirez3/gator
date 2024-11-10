# 🐊 GATOR
A blog aggreGATOR

## Requirements
Here are the project requirements for running the program.

### Go Installation
Install [Go](https://go.dev/doc/install).

### Postgres Installation
Install [Postgres](https://www.postgresql.org/).

### Config file
Add `~/.gatorconfig.json` to your home directory and populate with db config settings:
```json
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```
## Install Gator Executable
After cloning the repo, go to the project root and run
```bash
go install
```

## Run
From project root:
```bash
gator <command> [args...]
```

## Commands
- register <user>
  - Registers a new user
```bash
gator register homer
```

- login <user>
  - Logs in as a registered user
```bash
gator login bart
```

- addfeed <url> <title>
  - Adds & follows a feed
```bash
gator addfeed https://hnrss.org/newest hackernews
```

- agg <timestring>
  - Aggregates posts every given duration 
    - Valid time units are `"ns", "us" (or "µs"), "ms", "s", "m", "h"`
```bash
gator agg 12h
```
- browse <postsCount>
  - Displays the latest `postCount` posts saved from your feeds, in descending order.
```bash
gator browse 5
```
- follow <feedUrl>
  - Follows an existing feed
```bash
gator follow https://hnrss.org/newest
```
- unfollow <feedUrl>
  - Unfollows an existing feed
```bash
gator unfollow https://hnrss.org/newest
```
- following
  - Displays all feeds followed by current user
```bash
gator following
```
- users
  - Displays all users
```bash
gator users
```
- feeds
  - Displays all feeds
```bash
gator feeds
```
- reset
  - Clears all data (user, feed, follow and posts)
```bash
gator reset
```