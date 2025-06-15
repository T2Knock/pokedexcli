package main

import (
	"strings"
	"time"

	"github.com/T2Knock/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(config)
}
