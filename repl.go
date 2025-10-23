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

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	noWhitespace := strings.TrimSpace(text)
	toLower := strings.ToLower(noWhitespace)
	words := strings.Fields(toLower)
	return words
}

type cliCommand struct {
	name        string
	description string
	config      string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			config:      "",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 Pokemon world location areas",
			config:      "",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 Pokemon world location areas",
			config:      "",
			callback:    commandMap,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			config:      "",
			callback:    commandExit,
		},
	}
}
