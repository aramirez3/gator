package main

import (
	"context"
	"fmt"

	"github.com/aramirez3/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("unfollow requires a feed url")
	}
	url := cmd.arguments[0]
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return err
	}

	err = deleteFeedFollowRow(s, user, feed)
	if err != nil {
		return nil
	}

	return nil
}

func deleteFeedFollowRow(s *state, user database.User, feed database.Feed) error {
	params := database.DeleteFeedFollowsForUserAndFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err := s.db.DeleteFeedFollowsForUserAndFeed(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Printf("%s has unfollowed %s\n", user.Name, feed.Name)
	return nil
}
