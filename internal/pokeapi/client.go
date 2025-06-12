package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/T2Knock/pokedexcli/internal/pokecache"
)

func FetchLocationArea(url string, cache *pokecache.Cache) (LocationAreaResponse, error) {
	var result LocationAreaResponse

	if data, found := cache.Get(url); found {
		if err := json.Unmarshal(data, &result); err != nil {
			return LocationAreaResponse{}, fmt.Errorf("failed to parse JSON: %w", err)
		}
		return result, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to Get: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	cache.Add(url, data)

	if err = json.Unmarshal(data, &result); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return result, nil
}
