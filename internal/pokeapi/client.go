package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchLocationArea(url string) (LocationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to Get: %w", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	defer res.Body.Close()

	var result LocationAreaResponse

	err = json.Unmarshal(data, &result)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return result, nil
}
