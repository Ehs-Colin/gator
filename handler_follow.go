package main

import (
	"context"
	"fmt"

	"github.com/Ehs-Colin/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(state *State, cmd Command, user database.User, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	rssUrl := args[0]

	feed, err := state.db.GetFeedByUrl(context.Background(), rssUrl)
	if err != nil {
		return err
	}

	_, err = state.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Added %s to follow list for %s!\n", feed.Name, user.Name)
	printFeed(feed, user)
	return nil
}

func HandlerUnfollow(state *State, cmd Command, user database.User, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	rssUrl := args[0]
	feed, err := state.db.GetFeedByUrl(context.Background(), rssUrl)
	if err != nil {
		return err
	}
	err = state.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("%s unfollowed by %s\n", feed.Name, user.Name)
	return nil
}

func HandlerFollowing(state *State, cmd Command, user database.User, args ...string) error {
	feeds, err := state.db.GetFeedFollowersForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	fmt.Printf("%d Feed(s) followed by %s:\n", len(feeds), user.Name)
	for _, feed := range feeds {
		rssFeed := database.Feed{
			ID:        feed.ID,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			Name:      feed.Name,
			Url:       feed.Url,
			UserID:    feed.UserID,
		}
		printFeed(rssFeed, user)
	}
	return nil
}
