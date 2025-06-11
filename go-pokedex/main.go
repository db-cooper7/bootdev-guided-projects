package main

import (
	"time"

	"github.com/db-cooper7/bootdev-guided-projects/go-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(7 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startREPL(cfg)
}
