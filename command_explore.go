package main

import (
	"errors"
	"fmt"
)

// commandExplore ...
func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	locationResp, err := cfg.pokeapiClient.ListPokemonLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploreing %s...\n", cfg.commandArgs)
	fmt.Println("Found pokemon:")

	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
