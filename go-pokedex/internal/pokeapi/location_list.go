package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (LocationAreasResponse, error) {
	url := fmt.Sprintf("%s/location-area", baseURL)
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, exists := c.cache.Get(url); exists {
		locationRes := LocationAreasResponse{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreasResponse{}, errors.New("invalid request")
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationRes := LocationAreasResponse{}
	err = json.Unmarshal(dat, &locationRes)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locationRes, nil
}
