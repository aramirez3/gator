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
	fmt.Println("check args")
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("register requires a name")
	}
	user := cmd.arguments[0]

	ctx := context.Background()

	fmt.Println("check for existing user")

	existingUser, _ := s.db.GetUser(ctx, user)

	if existingUser.ID != uuid.Nil {
		fmt.Println("user exists")
		os.Exit(1)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Println("create new user")

	dbUser, err := s.db.CreateUser(ctx, params)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	fmt.Println("set cofig.current user")

	s.config.SetUser(dbUser.Name)

	fmt.Println("User was created!")
	fmt.Println(dbUser)
	return nil
}
