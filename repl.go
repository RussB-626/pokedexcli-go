package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]
		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	noWhitespace := strings.TrimSpace(text)
	toLower := strings.ToLower(noWhitespace)
	words := strings.Fields(toLower)
	return words
}
