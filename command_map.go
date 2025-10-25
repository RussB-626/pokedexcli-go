package main

import (
	"fmt"

	"github.com/russb-626/pokedexcli_go/internal/pokeapi"
)

func commandMap(config *Config) error {
	locationAreas, err := pokeapi.GetLocationAreas(config.Next)
	if err != nil {
		return err
	}

	if config.Next <= 20 {
		config.Back = 0
	} else {
		config.Back = config.Next - 20
	}

	if config.Next < locationAreas.Count {
		config.Next += 20
	}

	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}
