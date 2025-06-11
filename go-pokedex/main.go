package main

import (
	"time"

	"github.com/db-cooper7/bootdev-guided-projects/go-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(7*time.Second, 3*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startREPL(cfg)
}
