package main

import (
	"context"
	"fmt"
	"html"
)

func HandlerGetRSS(state *State, cmd Command, args ...string) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("Title: %s\n", html.UnescapeString(rssFeed.Channel.Title))
	fmt.Printf("Link: %s\n", rssFeed.Channel.Link)
	fmt.Printf("Description: %s\n", html.UnescapeString(rssFeed.Channel.Description))
	fmt.Println("Items:")
	for _, item := range rssFeed.Channel.Item {
		fmt.Println("Title: ", html.UnescapeString(item.Title))
		fmt.Println("Link: ", item.Link)
		fmt.Println("Description: ", html.UnescapeString(item.Description))
		fmt.Println("PubDate: ", html.UnescapeString(item.PubDate))
		fmt.Println("")
	}
	return nil
}
