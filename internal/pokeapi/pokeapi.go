package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/noxventura/pokedex-cli-go/internal/pokecache"
)

type pokeStructs interface {
	PokeJson | Encounter
}

func GetPokeData[T pokeStructs](url string, jsonStruct T) (T, error) {
	pokeCache := pokecache.NewCache(time.Second * 5)

	// cache get
	if item, ok := pokeCache.Get(url); ok {
		var currentJson T
		if err := json.Unmarshal(item, &currentJson); err != nil {
			var emptyType T
			return emptyType, fmt.Errorf("failed to unmarshal JSON: %w", err)
		}

		return currentJson, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		var emptyType T
		return emptyType, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		var emptyType T
		return emptyType, fmt.Errorf("failed to read response body: %w", err)
	}

	var currentJson T
	if err := json.Unmarshal(body, &currentJson); err != nil {
		var emptyType T
		return emptyType, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	// add to cache
	pokeCache.Add(url, body)
	return currentJson, nil
}

type PokeJson struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Encounter struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
