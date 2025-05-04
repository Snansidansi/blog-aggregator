package main

import (
	"context"
	"fmt"

	"github.com/snansidansi/blog-aggregator/internal/database"
)

func middleWareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("unable to get current user: %v", err)
		}

		return handler(s, cmd, user)
	}
}
