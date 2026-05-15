package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/somervillanthony/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	var locationAreaData pokeapi.LocationAreas

	cachedInfo, ok := cfg.cache.Get(url)
	if ok {
		err := json.Unmarshal(cachedInfo, &locationAreaData)
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
		err = json.Unmarshal(toBeCached, &locationAreaData)
		if err != nil {
			return fmt.Errorf("failed to decode json: %w", err)
		}
		cfg.cache.Add(url, toBeCached)
		fmt.Println("below is url")
		fmt.Println(url)
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

	var locationAreaData pokeapi.LocationAreas

	cachedInfo, ok := cfg.cache.Get(url)
	if ok {
		err := json.Unmarshal(cachedInfo, &locationAreaData)
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
		err = json.Unmarshal(toBeCached, &locationAreaData)
		if err != nil {
			return fmt.Errorf("failed to decode json: %w", err)
		}
		cfg.cache.Add(url, toBeCached)
	}

	cfg.Next = locationAreaData.Next
	cfg.Previous = locationAreaData.Previous

	for _, location := range locationAreaData.Results {
		fmt.Println(location.Name)
	}

	return nil
}
