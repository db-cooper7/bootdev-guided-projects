package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonSpecies(pokemonName string) (PokemonSpecies, error) {
	url := fmt.Sprintf("%s/pokemon-species/%s", baseURL, pokemonName)

	if val, ok := c.cache.Get(url); ok {
		pokemonSpeciesRes := PokemonSpecies{}
		err := json.Unmarshal(val, &pokemonSpeciesRes)
		if err != nil {
			return PokemonSpecies{}, err
		}

		return pokemonSpeciesRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonSpecies{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonSpecies{}, err
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return PokemonSpecies{}, errors.New("invalid pokemon")
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonSpecies{}, err
	}

	pokemonSpeciesRes := PokemonSpecies{}
	err = json.Unmarshal(dat, &pokemonSpeciesRes)
	if err != nil {
		return PokemonSpecies{}, err
	}

	c.cache.Add(url, dat)
	return pokemonSpeciesRes, nil
}
