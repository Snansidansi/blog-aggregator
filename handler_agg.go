package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/snansidansi/blog-aggregator/internal/database"
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

	err = savePostsToDB(s, rssFeed, nextFeedToFetch.ID)
	if err != nil {
		return fmt.Errorf("unable to save posts to the database: %v", err)
	}

	return nil
}

func savePostsToDB(s *state, rssFeed *RSSFeed, feedID uuid.UUID) error {
	for _, post := range rssFeed.Channel.Item {
		published_at := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, post.PubDate); err == nil {
			published_at = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       post.Title,
			Url:         post.Link,
			Description: post.Description,
			PublishedAt: published_at,
			FeedID:      feedID,
		})
		if err != nil {
			if pgErr, ok := err.(*pq.Error); ok {
				if pgErr.Code == "23505" {
					continue
				}
				return fmt.Errorf("unable to create new post: %v", err)
			}
		}
	}

	return nil
}
