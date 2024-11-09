package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("error getting users from db: %w", err)
	}

	if len(users) > 0 {
		fmt.Println("Users:")
		for _, user := range users {
			if user.Name == s.config.CurrentUserName {
				fmt.Printf("* %s (current)\n", user.Name)
			} else {
				fmt.Printf("* %s\n", user.Name)
			}
		}
	} else {
		fmt.Println("No registered users.")
	}
	return nil
}
