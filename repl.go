package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/somervillanthony/pokedexcli/internal/pokeapi"
	"github.com/somervillanthony/pokedexcli/internal/pokecache"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}
	cfg.pokedex = map[string]pokeapi.PokemonData{}
	cfg.cache = pokecache.NewCache(5 * time.Second)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		inputString := scanner.Text()

		inputString = strings.ToLower(inputString)
		inputSlice := strings.Fields(inputString)

		commands := getCommands()

		command, ok := commands[inputSlice[0]]
		if ok {
			err := command.callback(cfg, inputSlice[1:])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	lwrText := strings.ToLower(text)
	return strings.Fields(lwrText)
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args []string) error
	config      *config
}

type config struct {
	Next     *string
	Previous *string
	cache    *pokecache.Cache
	pokedex  map[string]pokeapi.PokemonData
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Find pokemon types in a location area name",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect an already caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display the names of all caught pokemon",
			callback:    commandPokedex,
		},
	}
}
