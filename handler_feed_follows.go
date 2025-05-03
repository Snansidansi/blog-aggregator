package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/snansidansi/blog-aggregator/internal/database"
)

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to get current user: %v", err)
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("unable to get feed for given url: %v", err)
	}

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	createdFeedFollows, err := s.db.CreateFeedFollow(context.Background(), feedFollow)
	if err != nil {
		return fmt.Errorf("unable to create new feed follow: %v", err)
	}

	for _, createdFeedFollow := range createdFeedFollows {
		fmt.Printf("Feedname: %s\n", createdFeedFollow.Feedname)
		fmt.Printf("Username: %s\n", createdFeedFollow.Username)
		fmt.Println("")
	}

	return nil
}

func handlerGetFollowedFeeds(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to get current user: %v", err)
	}

	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("unable to get feeds for current user: %v", err)
	}

	for _, followedFeed := range followedFeeds {
		fmt.Printf("* %s\n", followedFeed.Feedname)
	}

	return nil
}
