package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aramirez3/gator/internal/config"
)

func main() {

	config, err := config.Read()

	if err != nil {
		fmt.Println("error reading config file")
		os.Exit(1)
	}

	appState := state{
		config: &config,
	}

	cmds := commands{
		commands: make(map[string]commandHandler),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	loginCommand := command{
		name:      cmdName,
		arguments: cmdArgs,
	}

	err = cmds.run(&appState, loginCommand)
	if err != nil {
		log.Fatal(err)
	}
}
