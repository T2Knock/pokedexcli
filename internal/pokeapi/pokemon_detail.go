package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) PokemonDetail(name string) (Pokemon, error) {
	url := baseURL + "pokemon/" + name

	var result Pokemon

	if cache, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(cache, &result); err != nil {
			return Pokemon{}, err
		}

		return result, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error fetching pokemon: %v", name)
	}

	if resp.StatusCode != 200 {
		return Pokemon{}, fmt.Errorf("pokemon not found")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &result); err != nil {
		return Pokemon{}, err
	}

	return result, nil
}
