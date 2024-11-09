package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aramirez3/gator/internal/config"
	"github.com/aramirez3/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	config, err := config.Read()

	if err != nil {
		fmt.Println("error reading config file")
		os.Exit(1)
	}

	db, err := sql.Open("postgres", config.DBUrl)
	if err != nil {
		fmt.Printf("could not connect to database: %v\n", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	appState := state{
		config: &config,
		db:     dbQueries,
	}

	cmds := commands{
		commands: make(map[string]commandHandler),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", appState.middlewareLoggedIn(handlerAddfeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", appState.middlewareLoggedIn(handlerFollow))
	cmds.register("following", appState.middlewareLoggedIn(handlerFollowing))

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
