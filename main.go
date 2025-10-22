package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	noWhitespace := strings.TrimSpace(text)
	toLower := strings.ToLower(noWhitespace)
	words := strings.Fields(toLower)
	return words
}
