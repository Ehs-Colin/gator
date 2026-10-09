package main

import (
	"context"
	"fmt"

	"github.com/Ehs-Colin/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerGetUsers(state *State, cmd Command, args ...string) error {
	dbUsers, err := state.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error retrieving User list: %w", err)
	}
	for _, user := range dbUsers {
		current := ""
		if user.Name == state.config.CurrentUserName {
			current = "(current)"
		}
		fmt.Printf(" * %s %s\n", user.Name, current)
	}
	return nil
}

func HandlerReset(state *State, cmd Command, args ...string) error {
	err := state.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error reseting system: %w", err)
	}
	fmt.Println("System has been reset")
	return nil
}

func HandlerLogin(state *State, cmd Command, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := args[0]
	if state.config == nil {
		return fmt.Errorf("config is not initialized")
	}
	dbUser, err := state.db.GetUserByName(context.Background(), username)
	if err != nil {
		return err
	}
	err = state.config.SetUser(dbUser.Name)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	fmt.Printf("Logged in as: %v\n", username)
	return nil
}

func HandlerRegister(state *State, cmd Command, args ...string) error {
	if state.config == nil {
		return fmt.Errorf("config is not initialized")
	}
	if len(args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := args[0]
	// Add user to the database. Created_At and Updated_At are automatically set.
	dbUser, err := state.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:   uuid.New(),
		Name: username,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	err = state.config.SetUser(dbUser.Name)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	fmt.Printf("User %s created successfully!\n", dbUser.Name)
	printUser(dbUser)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
