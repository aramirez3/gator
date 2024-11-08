package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aramirez3/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("login requires a username")
	}

	name := cmd.arguments[0]

	user, exists := userExists(s, name)
	if !exists {
		fmt.Println("user exists")
		os.Exit(1)
	}
	s.config.SetUser(user.Name)
	fmt.Printf("Username has been set to: %s\n", user.Name)
	return nil
}

func userExists(s *state, name string) (database.User, bool) {
	ctx := context.Background()

	fmt.Println("check for existing user")

	user, _ := s.db.GetUser(ctx, name)

	if user.ID != uuid.Nil {
		return user, true
	}
	return user, false
}
