# Gator

A multi-line CLI for aggregating RSS feeds and viewing the posts

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database.
You can then install `gator` with:

```bash
go install https://github.com/db-cooper7/bootdev-guided-projects/go-rss-aggregator
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable",
  "current_user_name":"user_name_str",
  "limit":"max_amount_of_browse_posts"
}
```

Replace the values with your database connection string and the limit of browsed posts.

## Usage

Create a new user:

```bash
gator register <name>
```

Add a feed:

```bash
gator addfeed <name> <url>
```

Start the aggregator:

```bash
gator agg 30s
```

View the posts:

```bash
gator browse [limit]
```

There are a few other commands you'll need to know as well:

- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database