package main

import (
	"context"
	"fmt"
)

func handlerStartAggregator(s *state, cmd command) error {
	const url = "https://www.wagslane.dev/index.xml"

	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", rssFeed)
	return nil
}
