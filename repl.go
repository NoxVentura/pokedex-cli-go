package main

import (
	"bufio"
	"fmt"
	"internal/pokeapi"
	"os"
	"strings"
)

func StartRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		if isStopped {
			break
		}
		fmt.Print("pokedex > ")

		if !reader.Scan() {
			break
		}
		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Printf("command not found: %s\n\n", words[0])
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

var isStopped bool = false
var url string = ""
var pokeRes pokeapi.PokeJson = pokeapi.PokeJson{}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	words := strings.Fields(loweredText)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
	}
}
