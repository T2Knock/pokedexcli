package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/T2Knock/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			continue
		}

		input := scanner.Text()
		parseInput := cleanInput(input)
		command := parseInput[0]
		registerCommand, exist := getCommands()[command]

		if !exist {
			fmt.Println("Unknown command")
			continue
		}

		err := registerCommand.callback(config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the name of locations incrementally",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the name of locations decrementally",
			callback:    commandMapb,
		},
	}
}
