package handler

import (
	"fmt"

	"github.com/snansidansi/blog-aggregator/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	RegisteredCommands map[string]func(*state.State, Command) error
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	if callback, ok := c.RegisteredCommands[cmd.Name]; ok {
		err := callback(s, cmd)
		return err
	}

	return fmt.Errorf("Command does not exist: %s\n", cmd.Name)
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) error {
	if _, ok := c.RegisteredCommands[name]; ok {
		return fmt.Errorf("Command already exists: %s\n", name)
	}

	c.RegisteredCommands[name] = f
	return nil
}
