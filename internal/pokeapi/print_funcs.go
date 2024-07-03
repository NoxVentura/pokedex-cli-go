package pokeapi

import "fmt"

func PrintAreas(pokeRes *PokeJson) {
	areas := pokeRes.Results
	for _, value := range areas {
		fmt.Println(value.Name)
	}
}

func PrintEncounters(pokeRes *Encounter) {
	encounters := pokeRes.PokemonEncounters
	for _, value := range encounters {
		fmt.Println("- " + value.Pokemon.Name)
	}
}
