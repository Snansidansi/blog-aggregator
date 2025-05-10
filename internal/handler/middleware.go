package handler

import (
	"context"
	"fmt"

	"github.com/snansidansi/blog-aggregator/internal/database"
	"github.com/snansidansi/blog-aggregator/internal/state"
)

func MiddleWareLoggedIn(handler func(s *state.State, cmd Command, user database.User) error) func(*state.State, Command) error {
	return func(s *state.State, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("unable to get current user: %v", err)
		}

		return handler(s, cmd, user)
	}
}
