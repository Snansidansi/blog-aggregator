package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/snansidansi/blog-aggregator/internal/config"
	"github.com/snansidansi/blog-aggregator/internal/database"
)

func HandlerFollowFeed(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.Db.GetFeedByURL(context.Background(), url)
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

	createdFeedFollows, err := s.Db.CreateFeedFollow(context.Background(), feedFollow)
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

func HandlerUnfollowFeed(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feedUrl := cmd.Args[0]
	deleteFeedFollow := database.DeleteFeedFollowByUserIDAndFeedURLParams{
		UserID: user.ID,
		Url:    feedUrl,
	}

	err := s.Db.DeleteFeedFollowByUserIDAndFeedURL(context.Background(), deleteFeedFollow)
	if err != nil {
		return fmt.Errorf("unable to unfollow feed: %v", err)
	}

	fmt.Println("Successfuly unfollowed feed")
	return nil
}

func HandlerGetFollowedFeeds(s *config.State, cmd Command, user database.User) error {

	followedFeeds, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("unable to get feeds for current user: %v", err)
	}

	for _, followedFeed := range followedFeeds {
		fmt.Printf("* %s\n", followedFeed.Feedname)
	}

	return nil
}
