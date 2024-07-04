package main

import (
	"bufio"
	"fmt"
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

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		err := command.callback(args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

var isStopped bool = false

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	words := strings.Fields(loweredText)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
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
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explores the selected area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "To try to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Gives you information about the pokemon you've caught from the PokeDex",
			callback:    commandInspect,
		},
	}
}
