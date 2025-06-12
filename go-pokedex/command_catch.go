package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	pokemonSpecies, err := cfg.pokeapiClient.GetPokemonSpecies(args[0])
	if err != nil {
		return err
	}

	catchPercentage := (float64(pokemonSpecies.CaptureRate) / 255.0) * 100
	randomRoll := float64(rand.IntN(100))
	if randomRoll < catchPercentage {
		cfg.pokemonsCaught[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command")
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemon.Name)
	return nil
}
