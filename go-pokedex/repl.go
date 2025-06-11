package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/db-cooper7/bootdev-guided-projects/go-pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		cmd, exists := getCommands()[input[0]]
		if !exists {
			fmt.Println("invalid command")
			continue
		}
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations in the pokemon world",
			callback:    commandMapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the pokemon world",
			callback:    commandMapPrev,
		},
	}
}
