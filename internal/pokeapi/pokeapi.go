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
