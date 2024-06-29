package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokeJson struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetPokeData(url string) (PokeJson, error) {
	resp, err := http.Get(url)
	if err != nil {
		return PokeJson{}, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeJson{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var currentJson PokeJson
	if err := json.Unmarshal(body, &currentJson); err != nil {
		return PokeJson{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return currentJson, nil
}

