package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/db-cooper7/bootdev-guided-projects/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	if strings.TrimSpace(cmd.Args[0]) == "" {
		return fmt.Errorf("url cannot be empty")
	}

	feedID, err := s.db.GetFeedIDByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("could not get feed %w", err)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedID,
	})

	if err != nil {
		return fmt.Errorf("could not create feed follow: %w", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(feedFollowRow.UserName, feedFollowRow.FeedName)
	return nil
}

func printFeedFollow(userName, feedName string) {
	fmt.Printf("- User:			%s\n", userName)
	fmt.Printf("- Feed:			%s\n", feedName)
}

func handlerListFeedFollows(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s\n", user.Name)
	for _, feedfollow := range feedFollows {
		fmt.Printf("- %s\n", feedfollow.FeedName)
	}
	return nil
}
