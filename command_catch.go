package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/noxventura/pokedex-cli-go/internal/pokeapi"
)

func commandCatch(args ...string) error {
	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]

	pokemon, err := pokeapi.GetPokeData(url, pokeapi.Pokemon{})
	if err != nil {
		return errors.New("invalid pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	if isCaught(pokemon) {
		pokeDex := &appInstance.PokeDex
		(*pokeDex)[args[0]] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Printf("you may now inspect this pokemon using inspect\n")

	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func isCaught(pokemon pokeapi.Pokemon) bool {
	catchable := rand.Intn(pokemon.BaseExperience) < 40
	return catchable
}
