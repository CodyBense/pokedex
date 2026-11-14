package main

import (
	"errors"
	"fmt"
)

// commandInspect ...
func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("You must provide a pokemon's name")
	}

	name := args[0]

	pokemon := cfg.caughtPokemon[name]

	if pokemon.Name == "" {
		return errors.New("You don't have this pokemon")
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("\t- %s\n", t.Type.Name)
	}

	return nil
}
