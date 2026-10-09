package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Ehs-Colin/gator/internal/config"
	"github.com/Ehs-Colin/gator/internal/database"
	_ "github.com/lib/pq"
)

type State struct {
	config *config.Config
	db     *database.Queries
}

func main() {

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}
	state := &State{config: cfg}

	db, err := sql.Open("postgres", state.config.DbUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	state.db = dbQueries

	cmds := &Commands{commands: map[string]Command{}}

	cmds.Register("login", HandlerLogin)
	cmds.Register("register", HandlerRegister)
	cmds.Register("reset", HandlerReset)
	cmds.Register("users", HandlerGetUsers)
	cmds.Register("agg", HandlerGetRSS)
	cmds.Register("addfeed", middlewareLoggedIn(HandlerAddFeed))
	cmds.Register("feeds", HandlerFeeds)
	cmds.Register("follow", middlewareLoggedIn(HandlerFollow))
	cmds.Register("following", middlewareLoggedIn(HandlerFollowing))
	cmds.Register("unfollow", middlewareLoggedIn(HandlerUnfollow))
	cmds.Register("browse", middlewareLoggedIn(HandlerBrowse))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	commandName := os.Args[1]
	cmd, ok := cmds.Get(commandName)
	if !ok {
		log.Fatalf("command not found: %s", commandName)
	}
	args := os.Args[2:]
	if err := cmds.Run(state, cmd, args); err != nil {
		log.Fatal(err)
	}
}
