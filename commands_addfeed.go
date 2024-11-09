package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aramirez3/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddfeed(s *state, cmd command, user database.User) error {
	fmt.Println("Add feed command")
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("addfeed requires two arguments: name, url")
	}

	name := cmd.arguments[0]
	url := cmd.arguments[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)

	if err != nil {
		return err
	}

	fmt.Println("Feed created!")
	fmt.Printf("    - Name: %s\n", feed.Name)
	fmt.Printf("    - Url: %s\n", feed.Url)
	fmt.Printf("    - ID: %v\n", feed.ID)
	fmt.Printf("    - UserID: %v\n", feed.UserID)
	fmt.Printf("    - CreatedAt: %s\n", feed.CreatedAt.Format(time.RFC850))
	fmt.Printf("    - UpdatedAt: %s\n", feed.UpdatedAt.Format(time.RFC850))

	err = addFeedFollowRow(s, user.ID, feed.ID)
	if err != nil {
		return err
	}

	return nil
}
