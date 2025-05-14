package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			continue
		}

		input := scanner.Text()
		parseInput := cleanInput(input)

		fmt.Printf("Your command was: %v \n", parseInput[0])
	}
}
