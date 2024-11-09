package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aramirez3/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("follow requires a feed url")
	}

	url := cmd.arguments[0]
	feed, err := s.db.GetFeed(context.Background(), url)

	if err != nil {
		return err
	}

	err = addFeedFollowRow(s, user.ID, feed.ID)
	if err != nil {
		return err
	}

	return nil
}

func addFeedFollowRow(s *state, userId uuid.UUID, feedId uuid.UUID) error {
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userId,
		FeedID:    feedId,
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("* %s feed followed by %s\n", follow.FeedName, follow.UserName)

	return nil
}
