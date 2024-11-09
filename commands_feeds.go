package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)

	if err != nil {
		return fmt.Errorf("error getting feeds from db: %w", err)
	}

	if len(feeds) > 0 {
		fmt.Println("Feeds:")
		for _, feed := range feeds {
			fmt.Printf(" * %s\n", feed.Name)
			fmt.Printf("   - Url: %s\n", feed.Url)
			fmt.Printf("   - User: %s\n", feed.Username.String)
		}
	} else {
		fmt.Println("No feeds have been saved")
	}
	return nil
}
