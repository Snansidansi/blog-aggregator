package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/snansidansi/blog-aggregator/internal/config"
	"github.com/snansidansi/blog-aggregator/internal/database"
	"github.com/snansidansi/blog-aggregator/internal/handler"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	appState := config.State{
		Config: &conf,
	}

	db, err := sql.Open("postgres", appState.Config.DBURL)
	dbQueries := database.New(db)
	appState.Db = dbQueries

	commands := handler.Commands{
		RegisteredCommands: map[string]func(*config.State, handler.Command) error{},
	}

	commands.Register("login", handler.HandlerLogin)
	commands.Register("register", handler.HandlerRegister)
	commands.Register("reset", handler.HandlerResetDatabase)
	commands.Register("users", handler.HandlerGetUsers)
	commands.Register("agg", handler.HandlerStartAggregator)
	commands.Register("addfeed", handler.MiddleWareLoggedIn(handler.HandlerAddFeed))
	commands.Register("feeds", handler.HandlerGetFeeds)
	commands.Register("follow", handler.MiddleWareLoggedIn(handler.HandlerFollowFeed))
	commands.Register("following", handler.MiddleWareLoggedIn(handler.HandlerGetFollowedFeeds))
	commands.Register("unfollow", handler.MiddleWareLoggedIn(handler.HandlerUnfollowFeed))
	commands.Register("browse", handler.MiddleWareLoggedIn(handler.HandlerGetPosts))

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("expecting command name")
	}

	command := handler.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = commands.Run(&appState, command)
	if err != nil {
		log.Fatalln(err)
	}
}
