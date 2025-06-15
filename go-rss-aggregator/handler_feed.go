package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/db-cooper7/bootdev-guided-projects/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	if strings.TrimSpace(cmd.Args[0]) == "" || strings.TrimSpace(cmd.Args[1]) == "" {
		return fmt.Errorf("name and url cannot be empty")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		fmt.Printf("- Name: 		%s\n", feed.Name)
		fmt.Printf("- Url: 			%s\n", feed.Url)
		fmt.Printf("- Creator: 		%s\n", feed.Creator)
		fmt.Println()
	}
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("- ID:			%s\n", feed.ID)
	fmt.Printf("- Created:		%v\n", feed.CreatedAt)
	fmt.Printf("- Updated:		%v\n", feed.UpdatedAt)
	fmt.Printf("- Name:			%s\n", feed.Name)
	fmt.Printf("- URL:			%s\n", feed.Url)
	fmt.Printf("- UserID:		%s\n", feed.UserID)
	fmt.Println()
}
