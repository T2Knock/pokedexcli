package main

import (
	"errors"
	"fmt"
)

func commandMapf(config *config) error {
	locationsResp, error := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if error != nil {
		return error
	}

	config.nextLocationsURL = locationsResp.Next
	config.previousLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(config *config) error {
	if config.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := config.pokeapiClient.ListLocations(config.previousLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationResp.Next
	config.previousLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
