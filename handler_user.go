package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/snansidansi/blog-aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command 'register' expects a single argument")
	}

	username := cmd.args[0]

	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	createdUser, err := s.db.CreateUser(context.Background(), user)
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command 'login' expects a single argument")
	}

	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
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

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
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
