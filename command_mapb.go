package main

import (
	"fmt"

	"github.com/russb-626/pokedexcli_go/internal/pokeapi"
)

func commandMapb(config *Config) error {
	locationAreas, err := pokeapi.GetLocationAreas(config.Back)
	if err != nil {
		return err
	}

	if config.Back != 0 {
		config.Back -= 20
	}
	config.Next = config.Back + 20

	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}
