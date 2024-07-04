package main

import (
	"fmt"

	"github.com/noxventura/pokedex-cli-go/internal/pokeapi"
)

func commandInspect(args ...string) error {
	pokemon, ok := appInstance.PokeDex[args[0]]
	if !ok {
		return fmt.Errorf("%s : hasn't been caught yet", args[0])
	}

	pokemonPrinter(pokemon)
	return nil
}

func pokemonPrinter(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}
}
