package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *config, args []string) error {
	if len(config.pokedex) == 0 {
		return errors.New("pokedex is empty")
	}

	fmt.Println("Your Pokedex: ")
	for pokemon := range config.pokedex {
		fmt.Println(" - ", pokemon)
	}

	return nil
}
