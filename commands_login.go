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
		fmt.Println("user does not exist")
		os.Exit(1)
	}
	s.config.SetUser(user)
	fmt.Printf("Username has been set to: %s\n", user.Name)
	return nil
}

func userExists(s *state, name string) (database.User, bool) {
	user, _ := s.db.GetUser(context.Background(), name)

	if user.ID != uuid.Nil {
		return user, true
	}
	return user, false
}
