package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/snansidansi/blog-aggregator/internal/config"
	"github.com/snansidansi/blog-aggregator/internal/database"
)

func HandlerGetPosts(s *config.State, cmd Command, user database.User) error {
	postLimit := 2
	if len(cmd.Args) == 1 {
		convertedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("usage: %s <amount-of-posts>", cmd.Name)
		}

		postLimit = convertedLimit
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(postLimit),
	})
	if err != nil {
		return fmt.Errorf("unable to get posts from database: %v", err)
	}

	return printPosts(posts)
}

func printPosts(posts []database.GetPostsForUserRow) error {
	for _, post := range posts {
		fmt.Println("")
		fmt.Printf("* Title         : %s\n", post.Title)
		fmt.Printf("* Url           : %s\n", post.Url)
		fmt.Printf("* Description   : %s\n", post.Description)
		fmt.Printf("* Feed          : %s\n", post.FeedName)
		fmt.Printf("* Published at  : %s\n", post.PublishedAt.Time.Format("01.02.2006 Mon"))
	}

	return nil
}
