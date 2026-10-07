package main

import (
	"log"
	"os"

	"github.com/Ehs-Colin/gator/internal/config"
)

type State struct {
	config *config.Config
}

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}
	state := &State{config: cfg}
	cmds := &Commands{commands: map[string]Command{}}

	cmds.Register("login", Command{
		Name:   "login",
		args:   []string{},
		Action: HandlerLogin,
	})
	if len(os.Args) < 2 {
		log.Fatal("Usage: gator login <username>")
		return
	}
	cmd, ok := cmds.Get(os.Args[1])
	if !ok {
		log.Fatalf("command not found: login")
	}
	cmd.args = os.Args[2:]
	if err := cmds.Run(state, cmd); err != nil {
		log.Fatalf("failed to run command: %v", err)
	}
}
