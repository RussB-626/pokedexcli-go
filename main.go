package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			cleanInput := cleanInput(userInput)
			fmt.Println("Your command was: ", cleanInput[0])
		}
	}
}

func cleanInput(text string) []string {
	noWhitespace := strings.TrimSpace(text)
	toLower := strings.ToLower(noWhitespace)
	words := strings.Fields(toLower)
	return words
}
