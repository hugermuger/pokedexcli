package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, input ...string) error {
	pokemon, exists := cfg.pokedex[input[0]]
	if exists {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, x := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", x.Stat.Name, x.BaseStat)
		}
		fmt.Println("Types:")
		for _, i := range pokemon.Types {
			fmt.Printf("  - %v\n", i.Type.Name)
		}
	} else {
		return errors.New("you have not caught that pokemon")
	}
	return nil
}
