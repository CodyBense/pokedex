package main

import (
	"strings"
)

// cleanInput ...
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
