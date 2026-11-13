package main

import (
	"fmt"
	"os"
)

// commandExit ...
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
