package main

import (
	"errors"
	"fmt"
)

func commandPokedex(args ...string) error {
	if len(args) > 0 {
		return errors.New("too many arguments")
	}
	fmt.Println("Your Pokedex: ")
	for _, pokemon := range appInstance.PokeDex {
		fmt.Printf("  - %s \n", pokemon.Name)
	}
	return nil
}
