package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You have not yet caught any pokemon")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for pokemonName, _ := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}
