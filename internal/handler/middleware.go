package handler

import (
	"context"
	"fmt"

	"github.com/snansidansi/blog-aggregator/internal/config"
	"github.com/snansidansi/blog-aggregator/internal/database"
)

func MiddleWareLoggedIn(handler func(s *config.State, cmd Command, user database.User) error) func(*config.State, Command) error {
	return func(s *config.State, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("unable to get current user: %v", err)
		}

		return handler(s, cmd, user)
	}
}
