package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (LocationAreasResponse, error) {
	url := fmt.Sprintf("%s/location-area", baseURL)
	if pageUrl != nil {
		url = *pageUrl
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
