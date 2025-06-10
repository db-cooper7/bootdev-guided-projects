package main

import "strings"

func cleanInput(text string) []string {
	words := strings.ToLower(text)
	return strings.Fields(words)
}
