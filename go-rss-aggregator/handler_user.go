package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("could not set the username: %w", err)
	}
	fmt.Printf("The username has been set: %s\n", s.cfg.CurrentUserName)
	return nil
}
