package main

import "fmt"

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage: \n \n")

	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v \n", cmd.name, cmd.description)
	}

	return nil
}
