package main

import (
	"fmt"

	"github.com/aramirez3/gator/internal/config"
	"github.com/aramirez3/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commands map[string]commandHandler
}

type commandHandler func(*state, command) error

func (c *commands) register(name string, f commandHandler) {
	_, ok := c.commands[name]
	if !ok {
		c.commands[name] = f
	}
}

func (c *commands) run(s *state, cmd command) error {
	if cmdFunc, ok := c.commands[cmd.name]; ok {
		err := cmdFunc(s, cmd)
		if err != nil {
			return fmt.Errorf("error running command: %w", err)
		}
		return nil
	}
	return fmt.Errorf("could not find command")
}
