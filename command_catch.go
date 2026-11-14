package main

import (
	"fmt"
	"math"
	"math/rand"
)

// commandCatch ...
func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("You need to provide a pokemon name")
	}

	name := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	baseExperience := pokemon.BaseExperience
	catchRate := int(math.Round(float64(baseExperience) / 2))

	if rand.Intn(baseExperience) < catchRate {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
