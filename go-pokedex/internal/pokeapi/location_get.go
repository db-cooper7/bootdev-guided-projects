package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := fmt.Sprintf("%s/location-area/%s", baseURL, locationName)

	if val, exists := c.cache.Get(url); exists {
		var locationRes = Location{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return Location{}, err
		}
		return locationRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Location{}, errors.New("invalid locationName")
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationRes := Location{}
	err = json.Unmarshal(dat, &locationRes)
	if err != nil {
		return Location{}, nil
	}

	c.cache.Add(url, dat)
	return locationRes, nil
}
