package main

import (
	"context"
	"fmt"
	"time"

	"github.com/db-cooper7/bootdev-guided-projects/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("could not find the user: %w", err)
	}
	if err = s.cfg.SetUser(cmd.Args[0]); err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}

	fmt.Println("User set successfully")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	})

	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not get users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("- %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("- %v\n", user.Name)
	}

	return nil
}

func printUser(user database.User) {
	fmt.Printf("- ID:      %v\n", user.ID)
	fmt.Printf("- Name:    %v\n", user.Name)
}
