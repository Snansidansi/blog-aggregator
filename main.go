package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/snansidansi/blog-aggregator/internal/config"
	"github.com/snansidansi/blog-aggregator/internal/database"
)

type state struct {
	Config *config.Config
	db     *database.Queries
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	appState := state{
		Config: &conf,
	}

	db, err := sql.Open("postgres", appState.Config.DBURL)
	dbQueries := database.New(db)
	appState.db = dbQueries

	commands := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	commands.Register("login", handlerLogin)
	commands.Register("register", handlerRegister)
	commands.Register("reset", handlerResetDatabase)
	commands.Register("users", handlerGetUsers)
	commands.Register("agg", handlerStartAggregator)

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("expecting command name")
	}

	command := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = commands.Run(&appState, command)
	if err != nil {
		log.Fatalln(err)
	}

	conf, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	fmt.Printf("read config after changing current username: %+v\n", conf)
}
