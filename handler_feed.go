package main

import (
	"context"
	"fmt"

	"github.com/Ehs-Colin/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(state *State, cmd Command, args ...string) error {
	// 1> Validate command arguments for name and url
	if len(args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	feedName := args[0]
	feedUrl := args[1]
	// 2> Get the current user from db so user id can be written to new feed
	dbCurrentUser, err := state.db.GetUser(context.Background(), state.config.CurrentUserName)
	if err != nil {
		return err
	}
	// 3> Add the feed to the db.  Error if it already exists
	dbFeed, err := state.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   feedName,
		Url:    feedUrl,
		UserID: dbCurrentUser.ID,
	})
	if err != nil {
		return err
	}
	// 4> On success print out new feed info
	fmt.Println("Feed Id: ", dbFeed.ID)
	fmt.Println("Created At: ", dbFeed.CreatedAt)
	fmt.Println("Name: ", dbFeed.Name)
	fmt.Println("Url: ", dbFeed.Url)
	fmt.Println("User: ", state.config.CurrentUserName)
	return nil
}
