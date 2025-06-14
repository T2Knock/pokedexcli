package main

import "fmt"

func commandExplore(config *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("explore command requires a location name")
	}

	locationName := args[0]

	locationArea, err := config.pokeapiClient.LocationExplore(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", locationArea.Name)
	fmt.Println("Found Pokemon:")

	for _, pokemonEncounter := range locationArea.PokemonEncounters {
		fmt.Printf("- %v\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
