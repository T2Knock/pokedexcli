package pokeapi

import (
	"fmt"

	"github.com/T2Knock/pokedexcli/internal/pokecache"
)

type Config struct {
	NextURL     string
	PreviousURL string
}

func PrintNextLocation(config *Config, cache *pokecache.Cache) error {
	if config.NextURL == "" {
		fmt.Println("your're on the last page")

		return nil
	}

	locationAreas, err := FetchLocationArea(config.NextURL, cache)
	if err != nil {
		return err
	}

	if locationAreas.Next != nil {
		config.NextURL = *locationAreas.Next
	} else {
		config.NextURL = ""
	}

	if locationAreas.Previous != nil {
		config.PreviousURL = *locationAreas.Previous
	} else {
		config.PreviousURL = ""
	}

	if len(locationAreas.Results) > 0 {
		for _, result := range locationAreas.Results {
			fmt.Println(result.Name)
		}
	}

	return nil
}

func PrintPreviousLocation(config *Config, cache *pokecache.Cache) error {
	if config.PreviousURL == "" {
		fmt.Println("your're on the first page")

		return nil
	}

	locationAreas, err := FetchLocationArea(config.PreviousURL, cache)
	if err != nil {
		return err
	}

	if locationAreas.Next != nil {
		config.NextURL = *locationAreas.Next
	} else {
		config.NextURL = ""
	}

	if locationAreas.Previous != nil {
		config.PreviousURL = *locationAreas.Previous
	} else {
		config.PreviousURL = ""
	}

	if len(locationAreas.Results) > 0 {
		for _, result := range locationAreas.Results {
			fmt.Println(result.Name)
		}
	}

	return nil
}
