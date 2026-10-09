package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Ehs-Colin/gator/internal/database"
)

func HandlerBrowse(state *State, cmd Command, user database.User, args ...string) error {
	var err error
	limit := 2
	if len(args) > 0 {
		limit, err = strconv.Atoi(args[0])
		if err != nil {
			return err
		}
	}
	posts, err := state.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		printItem(post)
	}

	return nil
}

func printItem(post database.GetPostsForUserRow) {
	fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Url: %s\n", post.Url)
	fmt.Printf("Text: %s\n", post.Description.String)
	fmt.Println("============================================================")
}
