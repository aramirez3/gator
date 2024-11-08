package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aramirez3/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("register requires a name")
	}
	user := cmd.arguments[0]

	ctx := context.Background()

	existingUser, _ := s.db.GetUser(ctx, user)

	if existingUser.ID != uuid.Nil {
		os.Exit(1)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	dbUser, err := s.db.CreateUser(ctx, params)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	s.config.SetUser(dbUser.Name)

	fmt.Println("User was created!")
	fmt.Printf("    - ID: %v\n", dbUser.ID)
	fmt.Printf("    - Name: %s\n", dbUser.Name)
	fmt.Printf("    - Name: %s\n", dbUser.CreatedAt)
	fmt.Printf("    - Name: %s\n", dbUser.UpdatedAt)
	return nil
}
