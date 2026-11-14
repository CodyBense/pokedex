package main

import "fmt"

// commandPokedex ...
func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon)
	}

	return nil
}
