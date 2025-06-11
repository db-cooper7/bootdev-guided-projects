package main

import (
	"errors"
	"fmt"
)

func commandMapNext(cfg *config, args ...string) error {
	locationsRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsRes.Next
	cfg.prevLocationURL = locationsRes.Previous

	for _, location := range locationsRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapPrev(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're already on the first page")
	}

	locationsRes, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsRes.Next
	cfg.prevLocationURL = locationsRes.Previous

	for _, location := range locationsRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}
