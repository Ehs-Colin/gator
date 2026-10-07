package main

import "errors"

type Command struct {
	Name   string
	args   []string
	Action func(*State, Command) error
}

type Commands struct {
	commands map[string]Command
}

func (c *Commands) Register(name string, cmd Command) {
	c.commands[name] = cmd
}

func (c *Commands) Get(name string) (Command, bool) {
	cmd, ok := c.commands[name]
	return cmd, ok
}

func (c *Commands) Run(s *State, cmd Command) error {
	if cmd.Action == nil {
		return errors.New("command is not defined")
	}
	return cmd.Action(s, cmd)
}
