package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting users from table: %w", err)
	}
	fmt.Println("users table is reset.")

	err = s.db.DeleteAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting feed from table: %w", err)
	}
	fmt.Println("feed table is reset.")

	err = s.db.DeleteAllFeedFollows(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting from feed_follows table: %w", err)
	}

	err = s.db.DeleteAllPosts(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting posts: %w", err)
	}
	return nil
}
