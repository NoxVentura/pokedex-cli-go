package main

import "fmt"

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, value := range getCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	fmt.Println()
	return nil
}
