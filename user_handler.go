package main

import "fmt"

func handlerLogin(s *State, cmd Command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("Please enter a username: go run . %s <username>", cmd.name)
	}
	username := cmd.args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Error logging in user: %v", err)
	}
	fmt.Printf("User %v logged in\n", username)
	return nil
}
