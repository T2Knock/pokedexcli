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
	callback    func(config *config, args []string) error
}

type config struct {
	pokeapiClient        pokeapi.Client
	pokedex              map[string]pokeapi.Pokemon
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			return
		}

		input := scanner.Text()
		parseInput := cleanInput(input)

		if len(parseInput) == 0 {
			continue
		}

		command := parseInput[0]
		var args []string

		if len(parseInput) > 1 {
			args = parseInput[1:]
		}

		registerCommand, exist := getCommands()[command]

		if !exist {
			fmt.Println("Unknown command")
			continue
		}

		err := registerCommand.callback(config, args)
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
		"explore": {
			name:        "explore",
			description: "List all the pokemon located in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect caugth pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all the caught pokemon info",
			callback:    commandPokedex,
		},
	}
}
