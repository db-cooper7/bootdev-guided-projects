package main

import (
	"log"
	"os"

	"github.com/db-cooper7/bootdev-guided-projects/go-rss-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	runtimeState := state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <command_name> args ...string", os.Args[0])
		return
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmds.run(&runtimeState, cmd); err != nil {
		log.Fatal(err)
	}

}
