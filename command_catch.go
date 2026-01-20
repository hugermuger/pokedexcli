package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, input ...string) error {
	if len(input) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := input[0]
	locationsResp, err := cfg.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", locationsResp.Name)
	wurf := rand.Intn(255)
	if wurf >= locationsResp.BaseExperience {
		fmt.Printf("%v was caught!\n", locationsResp.Name)
		_, exists := cfg.pokedex[locationsResp.Name]
		if !exists {
			cfg.pokedex[locationsResp.Name] = locationsResp
		}
	} else {
		fmt.Printf("%v escaped!\n", locationsResp.Name)
	}

	return nil
}
