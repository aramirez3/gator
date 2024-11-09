package main

import (
	"context"
	"fmt"

	"github.com/aramirez3/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)

	if err != nil {
		return fmt.Errorf("error getting feeds from db: %w", err)
	}

	if len(feeds) > 0 {
		fmt.Printf("%s is following:\n", s.config.CurrentUserName)
		for _, feed := range feeds {
			fmt.Printf(" * %s\n", feed.FeedName)
		}
	} else {
		fmt.Printf("%s is not following any feeds\n", s.config.CurrentUserName)
	}
	return nil
}
