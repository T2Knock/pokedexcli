package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	var result LocationAreaResponse

	if err = json.Unmarshal(data, &result); err == nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return result, nil
}
