package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreas(offset int) (LocationAreas, error) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset)
	res, err := http.Get(fullUrl)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	var locationAreas LocationAreas
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return LocationAreas{}, err
	}
	return locationAreas, nil
}

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
