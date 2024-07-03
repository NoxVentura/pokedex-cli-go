package main

import (
	"errors"
	"fmt"

	"github.com/noxventura/pokedex-cli-go/internal/pokeapi"
)

type App struct {
	URL     string
	PokeRes *pokeapi.PokeJson
}

func (app *App) commandMap() error {
	if app.URL == "" {
		app.URL = "https://pokeapi.co/api/v2/location-area/"
	}

	pokeRes, err := pokeapi.GetPokeData(app.URL)
	if err != nil {
		return fmt.Errorf("failed to fetch PokeData: %w", err)
	}
	app.PokeRes = &pokeRes

	pokeapi.PrintAreas(pokeRes)
	app.URL = pokeRes.Next
	return nil
}

func (app *App) commandMapB() error {
	if app.PokeRes == nil || app.PokeRes.Previous == "" || app.URL == "" {
		return errors.New("no previous area to go back to")
	}
	app.URL = app.PokeRes.Previous

	pokeRes, err := pokeapi.GetPokeData(app.URL)
	if err != nil {
		return fmt.Errorf("failed to fetch PokeData: %w", err)
	}
	app.PokeRes = &pokeRes

	pokeapi.PrintAreas(pokeRes)
	return nil
}

var appInstance = &App{
	URL:     "https://pokeapi.co/api/v2/location-area/",
	PokeRes: nil,
}

func commandMap() error {
	return appInstance.commandMap()
}

func commandMapB() error {
	return appInstance.commandMapB()
}
