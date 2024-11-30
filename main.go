package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skarsden/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	programState := &State{cfg: &cfg}
	cmds := Commands{cmdNames: make(map[string]func(*State, Command) error)}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 3 {
		fmt.Println("not enough arguements, please enter: >go run . <command> <args[...]>")
		os.Exit(1)
	}

	commandName := os.Args[1]

	cmd := Command{name: commandName, args: os.Args[2:]}
	cmds.run(programState, cmd)
}

type State struct {
	cfg *config.Config
}
