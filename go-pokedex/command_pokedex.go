package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokemonsCaught) == 0 {
		return errors.New("no pokemon caught")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokemonsCaught {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
