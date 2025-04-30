package main

import (
	"fmt"
	"log"
	"os"

	"github.com/snansidansi/blog-aggregator/internal/config"
)

type state struct {
	Config *config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	fmt.Printf("Read config: %+v\n", conf)

	appState := state{
		Config: &conf,
	}

	commands := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	commands.Register("login", handlerLogin)

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
