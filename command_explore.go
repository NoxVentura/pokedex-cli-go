package main

import (
	"errors"

	"github.com/noxventura/pokedex-cli-go/internal/pokeapi"
)

func commandExplore(args ...string) error {
	if len(args) > 1 {
		return errors.New("too many arguments")
	}
	var encounter pokeapi.Encounter
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	pokeapi.GetPokeData(url, encounter)
}
