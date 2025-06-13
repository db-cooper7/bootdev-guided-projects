package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}

	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("could not reset users table: %w", err)
	}

	fmt.Println("Database reset successfully")
	return nil
}
