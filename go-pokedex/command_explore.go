package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location name provided")
	}
	locationRes, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("exploring %s...\n", locationRes.Name)
	fmt.Printf("Encountered Pokemon:\n")
	for _, encounter := range locationRes.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
