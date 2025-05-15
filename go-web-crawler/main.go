package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	normalizedURL, err := normalizeURL("https://www.boot.dev/lessons/98ac1f38-22dd-4682-b114-8638a0625567")
	if err != nil {
		return
	}
	fmt.Println(normalizedURL)
}
