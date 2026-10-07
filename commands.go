package main

import "errors"

type Command struct {
	Name   string
	Action func(*State, Command, ...string) error
}

type Commands struct {
	commands map[string]Command
}

func (c *Commands) Register(name string, f func(*State, Command, ...string) error) {
	cmd := Command{
		Name:   name,
		Action: f,
	}
	c.commands[name] = cmd
}

func (c *Commands) Get(name string) (Command, bool) {
	cmd, ok := c.commands[name]
	return cmd, ok
}

func (c *Commands) Run(s *State, cmd Command, args []string) error {
	if cmd.Action == nil {
		return errors.New("command is not defined")
	}
	return cmd.Action(s, cmd, args...)
}
