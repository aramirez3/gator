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
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Items       []RSSItem `xml:"item"`
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
		return fmt.Errorf("fetch requires a timeDuration string (e.g. 1s, 10m, 1h)")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("invalid time string: %w", err)
	}

	fmt.Printf("Collecting feed every %s\n", cmd.arguments[0])
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		fmt.Println("Run scrapeFeeds")
		scrapeFeeds(s)
	}
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

func getNextFeedToScrape(s *state) (database.Feed, error) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return database.Feed{}, fmt.Errorf("error getting feed: %w", err)
	}
	return feed, nil
}

func markFeedFetched(s *state, feed database.Feed) error {
	sqlTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	params := database.MarkFeedFetchedParams{
		ID:            feed.ID,
		LastFetchedAt: sqlTime,
	}

	err := s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return err
	}
	return nil
}

func scrapeFeeds(s *state) {
	feed, err := getNextFeedToScrape(s)
	if err != nil {
		fmt.Printf("Could not find feed to scrape: %v", err)
		return
	}

	fmt.Println("Found feed to scrape")
	scrapeFeed(s, feed)
}

func scrapeFeed(s *state, feed database.Feed) error {
	err := markFeedFetched(s, feed)
	if err != nil {
		return err
	}

	feedData, err := fetchFeed(feed.Url)
	if err != nil {
		return fmt.Errorf("could not retrieve feed data: %w", err)
	}

	fmt.Printf("Fetching items for %s (%v posts found):\n", feedData.Channel.Title, len(feedData.Channel.Items))
	for _, item := range feedData.Channel.Items {
		postPubDate, _ := time.Parse(time.Layout, item.PubDate)
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: postPubDate,
			FeedID:      feed.ID,
		}
		_, err := s.db.CreatePost(context.Background(), params)
		if err != nil {
			fmt.Printf(" x error: %v\n", err)
		}
		fmt.Printf(" * added: %s\n", item.Title)
	}
	return nil
}
