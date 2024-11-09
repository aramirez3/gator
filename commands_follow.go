package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aramirez3/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("follow requires a feed url")
	}

	ctx := context.Background()
	url := cmd.arguments[0]
	feed, err := s.db.GetFeed(ctx, url)

	if err != nil {
		return err
	}

	err = addFeedFollowRow(s, feed.ID)
	if err != nil {
		return err
	}

	return nil
}

func addFeedFollowRow(s *state, feedId uuid.UUID) error {
	ctx := context.Background()

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    s.config.CurrentUserId,
		FeedID:    feedId,
	}

	follow, err := s.db.CreateFeedFollow(ctx, params)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("* %s feed followed by %s\n", follow.FeedName, follow.UserName)

	return nil
}
