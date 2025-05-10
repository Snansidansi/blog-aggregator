package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/snansidansi/blog-aggregator/internal/config"
	"github.com/snansidansi/blog-aggregator/internal/database"
)

func HandlerRegister(s *config.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("command 'register' expects a single argument")
	}

	username := cmd.Args[0]

	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	createdUser, err := s.Db.CreateUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("user already exists: %s", username)
	}

	err = s.Config.SetUser(username)
	if err != nil {
		return fmt.Errorf("unable to set current user: %v\n", err)
	}

	fmt.Println("User was registered successfuly:")
	printUser(createdUser)

	return nil
}

func printUser(createdUser database.User) {
	fmt.Printf("%-10s: %v\n", "Username", createdUser.Name)
	fmt.Printf("%-10s: %v\n", "ID", createdUser.ID)
}

func HandlerLogin(s *config.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("command 'login' expects a single argument")
	}

	username := cmd.Args[0]
	_, err := s.Db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user does not exist: %s", username)
	}

	err = s.Config.SetUser(username)
	if err != nil {
		return fmt.Errorf("unable to set current user: %v\n", err)
	}

	fmt.Printf("Username set to: %s\n", username)
	return nil
}

func HandlerGetUsers(s *config.State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to fetch users: %v", err)
	}

	for _, user := range users {
		current := ""
		if user.Name == s.Config.CurrentUserName {
			current = " (current)"
		}

		fmt.Printf("* %s%s\n", user.Name, current)
	}
	return nil
}
