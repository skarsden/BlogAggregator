package main

import "fmt"

type Command struct {
	name string
	args []string
}

type Commands struct {
	cmdNames map[string]func(*State, Command) error
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.cmdNames[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
	f, ok := c.cmdNames[cmd.name]
	if !ok {
		return fmt.Errorf("Command %v does not exist", cmd.name)
	}

	f(s, cmd)
	return nil
}
