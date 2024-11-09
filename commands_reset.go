package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()

	err := s.db.DeleteAllUsers(ctx)
	if err != nil {
		return fmt.Errorf("error deleting users from table: %w", err)
	}
	fmt.Println("users table is reset.")

	err = s.db.DeleteAllFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error deleting feed from table: %w", err)
	}
	fmt.Println("feed table is reset.")

	err = s.db.DeleteAllFeedFollows(ctx)
	if err != nil {
		return fmt.Errorf("error deleting from feed_follows table: %w", err)
	}

	return nil
}
