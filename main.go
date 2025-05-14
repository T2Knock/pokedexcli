package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}

func commandHelp(commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage: \n \n")

	for _, command := range commands {
		fmt.Printf("%v: %v \n", command.name, command.description)
	}

	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var commands map[string]cliCommand

	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp(commands) },
		},
	}

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			continue
		}

		input := scanner.Text()
		parseInput := cleanInput(input)
		command := parseInput[0]
		registerCommand, exist := commands[command]

		if !exist {
			fmt.Println("Unknown command")
			continue
		}

		err := registerCommand.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
