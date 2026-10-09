package main

import (
	"context"
	"fmt"
	"html"
	"time"
)

func HandlerGetRSS(state *State, cmd Command, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: %s <time>", cmd.Name)
	}
	timeBetweenRequests, err := time.ParseDuration(args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(state)
	}
	// rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("Title: %s\n", html.UnescapeString(rssFeed.Channel.Title))
	// fmt.Printf("Link: %s\n", rssFeed.Channel.Link)
	// fmt.Printf("Description: %s\n", html.UnescapeString(rssFeed.Channel.Description))
	// fmt.Println("Items:")
	// for _, item := range rssFeed.Channel.Item {
	// 	fmt.Println("Title: ", html.UnescapeString(item.Title))
	// 	fmt.Println("Link: ", item.Link)
	// 	fmt.Println("Description: ", html.UnescapeString(item.Description))
	// 	fmt.Println("PubDate: ", html.UnescapeString(item.PubDate))
	// 	fmt.Println("")
	// }
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
	fmt.Printf("RSS Feed: %s\n", html.UnescapeString(rssFeed.Channel.Title))
	fmt.Println("Items:")
	for _, item := range rssFeed.Channel.Item {
		fmt.Println("Title: ", html.UnescapeString(item.Title))
		fmt.Println("========================================")
	}
	return nil
}
