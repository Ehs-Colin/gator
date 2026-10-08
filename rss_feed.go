package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// 1> Create new GET request
	feedRequest, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	// 2> Do the request
	feedRequest.Header.Add("User-Agent", "gator")
	resp, err := http.DefaultClient.Do(feedRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 3> Read the request into dat
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// 4> Unmarshal XML into usable struct
	rssXml := &RSSFeed{}
	err = xml.Unmarshal(dat, rssXml)
	if err != nil {
		return nil, err
	}
	// 5> Update fields to correct strings from html
	rssXml.Channel.Title = html.UnescapeString(rssXml.Channel.Title)
	rssXml.Channel.Description = html.UnescapeString(rssXml.Channel.Description)
	for i, item := range rssXml.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssXml.Channel.Item[i] = item
	}
	// 6> Return the xml feed
	return rssXml, err
}
