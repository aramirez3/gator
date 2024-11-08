package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func handlerAggs(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("register requires a name")
	}

	feedURL := cmd.arguments[0]
	ctx := context.Background()
	rssFeed, err := fetchFeed(ctx, feedURL)
	if err != nil {
		return err
	}
	fmt.Println(rssFeed)
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	feed := RSSFeed{}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err != nil {
		return &feed, err
	}
	req.Header.Add("User-Agent", "gator")

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return &feed, err
	}

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return &feed, err
	}

	err = xml.Unmarshal(byteData, &feed)
	if err != nil {
		return &feed, err
	}

	return &feed, err
}
