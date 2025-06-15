package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("catch command requires a pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	pokemon, err := config.pokeapiClient.PokemonDetail(pokemonName)
	if err != nil {
		return err
	}

	caught := attemptCatch(pokemon.BaseExperience)
	if caught {
		fmt.Printf("%v was caught! \n", pokemonName)
		config.pokedex[pokemonName] = pokemon
		return nil
	}

	fmt.Println("You missed the Pokémon!")

	return nil
}

func calculateCatchRate(baseExperience int) int {
	maxBaseExperience := 608

	ratio := float64(baseExperience) / float64(maxBaseExperience)
	catchRate := 100 - int(ratio*95)

	return max(10, catchRate)
}

func attemptCatch(baseExperience int) bool {
	catchChance := calculateCatchRate(baseExperience)
	randomRoll := randomInt(1, 100)

	return randomRoll < catchChance
}

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
