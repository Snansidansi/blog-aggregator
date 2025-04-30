package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) Run(s *state, cmd command) error {
	if callback, ok := c.registeredCommands[cmd.name]; ok {
		err := callback(s, cmd)
		return err
	}

	return fmt.Errorf("Command does not exist: %s\n", cmd.name)
}

func (c *commands) Register(name string, f func(*state, command) error) error {
	if _, ok := c.registeredCommands[name]; ok {
		return fmt.Errorf("Command already exists: %s\n", name)
	}

	c.registeredCommands[name] = f
	return nil
}
