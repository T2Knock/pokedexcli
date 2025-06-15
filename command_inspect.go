package main

import "fmt"

func commandInspect(config *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("inspect command requires a pokemon name")
	}

	pokemonName := args[0]
	pokemon, exist := config.pokedex[pokemonName]
	if !exist {
		fmt.Printf("No information on %v yet\n", pokemonName)
		return nil
	}

	fmt.Printf("Name %v\n Height: %v\n Weight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("\t- %v\n", pokemonType.Type.Name)
	}

	return nil
}
