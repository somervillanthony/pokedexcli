package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"

	"github.com/somervillanthony/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("catch needs at least one additional parameter")
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]

	var pokemonData pokeapi.PokemonData
	cachedInfo, ok := cfg.cache.Get(url)
	if ok {
		err := json.Unmarshal(cachedInfo, &pokemonData)
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
		err = json.Unmarshal(toBeCached, &pokemonData)
		if err != nil {
			return fmt.Errorf("failed to decode json: %w", err)
		}
		cfg.cache.Add(url, toBeCached)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	randCatchChance := rand.IntN(100)
	catchThreshold := 100 - (pokemonData.BaseExperience / 5)

	if catchThreshold <= 49 {
		catchThreshold = 50
	}
	if randCatchChance < catchThreshold {
		fmt.Printf("%s was caught!\n", args[0])
		cfg.pokedex[args[0]] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", args[0])
	}

	return nil
}
