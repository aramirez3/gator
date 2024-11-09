package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/aramirez3/gator/internal/database"
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

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("fetch requires a feedUrl")
	}

	rssFeed, err := fetchFeed(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Println(rssFeed)
	return nil
}

func fetchFeed(feedURL string) (*RSSFeed, error) {
	feed := RSSFeed{}
	req, err := http.NewRequestWithContext(context.Background(), "GET", feedURL, nil)

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

	unescapeFeed(&feed)

	return &feed, err
}

func unescapeFeed(rssFeed *RSSFeed) *RSSFeed {
	title := html.UnescapeString(rssFeed.Channel.Title)
	description := html.UnescapeString(rssFeed.Channel.Description)
	rssFeed.Channel.Title = title
	rssFeed.Channel.Description = description
	return rssFeed
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}

	sqlTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	params := database.MarkFeedFetchedParams{
		ID:            nextFeed.ID,
		LastFetchedAt: sqlTime,
	}
	err = s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return err
	}
	return nil
}
