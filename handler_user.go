package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command 'login' expects a single argument")
	}

	username := cmd.args[0]
	err := s.Config.SetUser(username)
	if err != nil {
		return fmt.Errorf("Unable to set new user: %v\n", err)
	}

	fmt.Printf("Username set to: %s\n", username)
	return nil
}
