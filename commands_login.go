package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("login requires a username")
	}
	name := cmd.arguments[0]
	s.config.SetUser(name)
	fmt.Printf("Username has been set to: %s\n", name)
	return nil
}
