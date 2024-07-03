package main

import (
	"errors"
	"fmt"

	"github.com/noxventura/pokedex-cli-go/internal/pokeapi"
)

func commandExplore(args ...string) error {
	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	var encounter pokeapi.Encounter
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

	encounter, err := pokeapi.GetPokeData(url, encounter)
	if err != nil {
		return errors.New("invalid pokemon area")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon: ")
	pokeapi.PrintEncounters(&encounter)

	return nil
}
