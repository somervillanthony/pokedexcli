package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/somervillanthony/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("explore needs at least one additional parameter\n")
	}
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

	fmt.Printf("Exploring %s...\n", args[0])

	var exploreData pokeapi.LocationAreaExplore
	cachedInfo, ok := cfg.cache.Get(url)
	if ok {
		err := json.Unmarshal(cachedInfo, &exploreData)
		if err != nil {
			return fmt.Errorf("Failed to Unmarshal cached json info: %w", err)
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("HTTP GET failed: %w", err)
		}
		defer res.Body.Close()
		toBeCached, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("io readall to get json data to cache failed: %w", err)
		}
		err = json.Unmarshal(toBeCached, &exploreData)
		if err != nil {
			return fmt.Errorf("failed to decode json: %w", err)
		}
		cfg.cache.Add(url, toBeCached)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range exploreData.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
