package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/snansidansi/blog-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	if s.Config.CurrentUserName == "" {
		return fmt.Errorf("no user is currently logged in")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to fetch current user: %v", err)
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    currentUser.ID,
	}

	createdFeed, err := s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		return fmt.Errorf("unable to create new feed: %v", err)
	}

	fmt.Println("Created feed successfuly:")
	printFeed(createdFeed)
	fmt.Println("")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:      %s\n", feed.ID)
	fmt.Printf("* Created: %v\n", feed.CreatedAt)
	fmt.Printf("* Updated: %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:    %s\n", feed.Name)
	fmt.Printf("* URL:     %s\n", feed.Url)
	fmt.Printf("* UserID:  %s\n", feed.UserID)
}
