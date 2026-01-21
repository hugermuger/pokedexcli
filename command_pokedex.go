package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, input ...string) error {
	if len(cfg.pokedex) > 0 {
		fmt.Println("Your Pokedex:")
		for _, i := range cfg.pokedex {
			fmt.Printf(" - %v\n", i.Name)
		}
	} else {
		return errors.New("Your Pokedex is empty!")
	}
	return nil
}
