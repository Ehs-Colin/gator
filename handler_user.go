package main

import "fmt"

func HandlerLogin(state *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.args[0]
	if state.config == nil {
		return fmt.Errorf("config is not initialized")
	}
	err := state.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	fmt.Printf("Logged in as: %v\n", username)
	return nil
}
