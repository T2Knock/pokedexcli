package pokeapi

import "fmt"

type Config struct {
	NextURL     string
	PreviousURL string
}

func PrintNextLocation(config *Config) error {
	if config.NextURL == "" {
		fmt.Println("your're on the last page")

		return nil
	}

	locationAreas, err := FetchLocationArea(config.NextURL)
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

func PrintPreviousLocation(config *Config) error {
	if config.PreviousURL == "" {
		fmt.Println("your're on the first page")

		return nil
	}

	locationAreas, err := FetchLocationArea(config.PreviousURL)
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
