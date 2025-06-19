package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/db-cooper7/bootdev-guided-projects/go-rss-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit, err := strconv.Atoi(s.cfg.DefaultLimit)
	if err != nil {
		return fmt.Errorf("could not get the default limit config variable: %w", err)
	}
	if len(cmd.Args) == 1 {
		wantedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}

		if wantedLimit < limit {
			limit = wantedLimit
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return fmt.Errorf("could not get posts for user: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("Date: %s\n", post.PublishedAt.Time.Format("Mon Jan 2"))
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("	%v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
	}

	return nil
}
