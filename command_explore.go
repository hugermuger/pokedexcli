package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, input ...string) error {
	if len(input) != 1 {
		return errors.New("you must provide a location name")
	}

	name := input[0]
	locationsResp, err := cfg.pokeapiClient.ListPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", locationsResp.Name)
	fmt.Println("Found Pokemon:")
	for _, loc := range locationsResp.PokemonEncounters {
		fmt.Printf(" - %v\n", loc.Pokemon.Name)
	}
	return nil
}
