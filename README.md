# GATOR – CLI RSS

A Boot.dev project for practicing Go with SQL by building a command-line RSS reader and aggregator.

## Installation

### Requirements

* Go installed
* PostgreSQL installed

### Install the CLI

`go install github.com/Tinotsu/gator`

## Usage

Run any command with:
`gator <command> [arguments]`

## Commands

### User Management

* `register [name]`
  Register a new user.

* `login [name]`
  Log in as an existing user.

* `users`
  List all registered users.

* `reset`
  Reset the database (destructive).

### Feeds

* `addfeed [name] [url]`
  Add a new RSS feed to the database (requires login).

* `feeds`
  List all available feeds.

### Following

* `follow [feed_name]`
  Follow a feed (requires login).

* `following`
  List the feeds you're following (requires login).

* `unfollow [feed_name]`
  Unfollow a feed (requires login).

### Reading

* `browse [limit]`
  Browse recent posts from followed feeds (requires login).

### Aggregation

* `agg`
  Run the RSS aggregator to fetch and store new posts.

## Configuration

Ensure your database connection URL is set in your configuration file or environment so the application can connect to PostgreSQL. The `main.go` uses the `DBURL` field from the config.

## Notes

* Some commands require that you are logged in (middleware enforces this).
* The `main.go` registers additional commands beyond those in earlier README drafts; include them here for completeness.

