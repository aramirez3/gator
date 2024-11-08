package main

import "fmt"

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("register requires a name")
	}
	user := cmd.arguments[0]
	fmt.Println("Register command")
	fmt.Printf("Create user in the database: %s\n", user)
	return nil
}
