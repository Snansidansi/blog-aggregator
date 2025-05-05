package main

import (
	"context"
	"fmt"
	"time"
)

func handlerStartAggregator(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time-between-reqs (e.g. 1s, 1m, 1h, 1h10m, ...)>", cmd.name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %v\nExample durations: 1s, 1m, 1h, 1h10m, ...", err)
	}

	fmt.Printf("Collecting feeds every: %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
		fmt.Println("")
	}
}

func scrapeFeeds(s *state) error {
	nextFeedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("unable to next feed to fetch: %v", err)
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeedToFetch.Url)
	if err != nil {
		return fmt.Errorf("unable to fetch the next feed: %v", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeedToFetch.ID)
	if err != nil {
		return fmt.Errorf("unable to march next feed as fetched: %v", err)
	}

	fmt.Printf("%d titles for %s:\n", len(rssFeed.Channel.Item), nextFeedToFetch.Name)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}

	return nil
}
