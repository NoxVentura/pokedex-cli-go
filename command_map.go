package main

import (
	"errors"
	"fmt"

	"github.com/noxventura/pokedex-cli-go/internal/pokeapi"
)

type App struct {
	URL     string
	PokeRes *pokeapi.PokeJson
	PokeDex map[string]pokeapi.Pokemon
}

func (app *App) commandMap(args ...string) error {
	if app.URL == "" {
		app.URL = "https://pokeapi.co/api/v2/location-area/"
	}
	pokeRes, err := pokeapi.GetPokeData(app.URL, pokeapi.PokeJson{})
	if err != nil {
		return fmt.Errorf("failed to fetch PokeData: %w", err)
	}
	app.PokeRes = &pokeRes

	pokeapi.PrintAreas(&pokeRes)
	app.URL = pokeRes.Next
	return nil
}

func (app *App) commandMapB(args ...string) error {
	if app.PokeRes == nil || app.PokeRes.Previous == "" || app.URL == "" {
		return errors.New("no previous area to go back to")
	}
	app.URL = app.PokeRes.Previous

	pokeRes, err := pokeapi.GetPokeData(app.URL, pokeapi.PokeJson{})
	if err != nil {
		return fmt.Errorf("failed to fetch PokeData: %w", err)
	}
	app.PokeRes = &pokeRes

	pokeapi.PrintAreas(&pokeRes)
	return nil
}

var appInstance = &App{
	URL:     "https://pokeapi.co/api/v2/location-area/",
	PokeRes: nil,
	PokeDex: make(map[string]pokeapi.Pokemon),
}

func commandMap(args ...string) error {
	return appInstance.commandMap()
}

func commandMapB(args ...string) error {
	return appInstance.commandMapB()
}
