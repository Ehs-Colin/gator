package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Ehs-Colin/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerGetRSS(state *State, cmd Command, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: %s <time>", cmd.Name)
	}
	timeBetweenRequests, err := time.ParseDuration(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s...", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(state)
	}
	return nil
}

func scrapeFeeds(state *State) error {
	// 1> Get next feed from database.  Next feed is oldest (or null) last fetched date
	dbFeed, err := state.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	// 2> Mark the feed as fetched so it goes to the back of the line
	_, err = state.db.MarkFeedFetched(context.Background(), dbFeed.ID)
	if err != nil {
		return err
	}
	// 3> Reed xml from feed url and print info
	rssFeed, err := fetchFeed(context.Background(), dbFeed.Url)
	if err != nil {
		return err
	}
	for _, item := range rssFeed.Channel.Item {
		SavePost(state, item, dbFeed)
	}
	return nil
}

func SavePost(state *State, item RSSItem, feed database.Feed) error {
	itemDescription := sql.NullString{
		String: item.Description,
		Valid:  len(item.Description) > 0,
	}
	publishedAt := sql.NullTime{}
	if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
		publishedAt = sql.NullTime{
			Time:  t,
			Valid: true,
		}
	}

	_, err := state.db.CreatePost(context.Background(), database.CreatePostParams{
		ID:          uuid.New(),
		Title:       item.Title,
		Url:         item.Link,
		Description: itemDescription,
		PublishedAt: publishedAt,
		FeedID:      feed.ID,
	})
	if err != nil {
		if !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			log.Printf("Couldn't create post: %v", err)
		}
	}

	return nil
}
