package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		inputString := scanner.Text()

		inputString = strings.ToLower(inputString)
		inputSlice := strings.Fields(inputString)

		commands := getCommands()

		command, ok := commands[inputSlice[0]]
		if ok {
			err := command.callback(cfg)
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
	callback    func(cfg *config) error
	config      *config
}

type config struct {
	Next     *string
	Previous *string
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
	}
}
