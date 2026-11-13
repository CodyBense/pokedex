package main

import (
	"errors"
	"fmt"
)

// commandMap ...
func commandMap(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationResp.Next
	cfg.previousLocationsUrl = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

// commandMapb ...
func commandMapb(cfg *config) error {
	if cfg.previousLocationsUrl == nil {
		return errors.New("You're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationResp.Next
	cfg.previousLocationsUrl = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
