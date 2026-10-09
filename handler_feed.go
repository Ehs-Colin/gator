package main

import (
	"context"
	"fmt"

	"github.com/Ehs-Colin/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(state *State, cmd Command, user database.User, args ...string) error {
	// 1> Validate command arguments for name and url
	if len(args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	feedName := args[0]
	feedUrl := args[1]
	// 2> Get user moved to middleware
	// 3> Add the feed to the db.  Error if it already exists
	dbFeed, err := state.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   feedName,
		Url:    feedUrl,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	_, err = state.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		FeedID: dbFeed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	// 4> On success print out new feed info
	fmt.Println("Feed created successfully:")
	printFeed(dbFeed, user)
	return nil
}

func printFeed(feed database.Feed, userName database.User) {
	fmt.Println("Feed Id:    ", feed.ID)
	fmt.Println("Created At: ", feed.CreatedAt)
	fmt.Println("Name:       ", feed.Name)
	fmt.Println("Url:        ", feed.Url)
	fmt.Println("User:       ", userName.Name)
	fmt.Println("=====================================")
}

func HandlerFeeds(state *State, cmd Command, args ...string) error {
	dbFeeds, err := state.db.GetAllFeeds(context.Background())
	if err != nil {
		return nil
	}
	if len(dbFeeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}
	fmt.Printf("Found %d feeds:\n", len(dbFeeds))
	for _, feed := range dbFeeds {
		user, err := state.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Couldn't get user: %w", err)
		}
		printFeed(feed, user)
	}

	return nil
}
