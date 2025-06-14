package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) LocationExplore(name string) (LocationAreaDetail, error) {
	url := baseURL + "location-area/" + name

	var result LocationAreaDetail

	if cache, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(cache, &result); err != nil {
			return LocationAreaDetail{}, err
		}

		return result, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaDetail{}, fmt.Errorf("error fetching location: %v", name)
	}

	if resp.StatusCode != 200 {
		return LocationAreaDetail{}, fmt.Errorf("location area not found")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &result); err != nil {
		return LocationAreaDetail{}, err
	}

	return result, nil
}
