package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CodyBense/pokedex/interal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	commands             map[string]cliCommand
	commandArgs          string
	pokeapiClient        pokeapi.Client
	nextLocationsUrl     *string
	previousLocationsUrl *string
	caughtPokemon        map[string]pokeapi.Pokemon
}

// getCommand ...
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
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all the pokemon in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Describe a pokemon in your Pokedex",
			callback:    commandInspect,
		},
	}

}

// cleanInput ...
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// startRepl ...
func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := cfg.commands[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}
	}
}
