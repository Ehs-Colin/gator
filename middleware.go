package main

import (
	"context"

	"github.com/Ehs-Colin/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User, args ...string) error) func(*State, Command, ...string) error {
	return func(state *State, cmd Command, args ...string) error {
		user, err := state.db.GetUserByName(context.Background(), state.config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(state, cmd, user, args...)
	}
}
