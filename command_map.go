package main

import (
	"fmt"

	"github.com/somervillanthony/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	locationAreaData, err := pokeapi.UnmarshalJson(url)
	if err != nil {
		return err
	}

	cfg.Next = locationAreaData.Next
	cfg.Previous = locationAreaData.Previous

	for _, location := range locationAreaData.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Previous != nil {
		url = *cfg.Previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
	locationAreaData, err := pokeapi.UnmarshalJson(url)
	if err != nil {
		return err
	}

	cfg.Next = locationAreaData.Next
	cfg.Previous = locationAreaData.Previous

	for _, location := range locationAreaData.Results {
		fmt.Println(location.Name)
	}

	return nil
}
