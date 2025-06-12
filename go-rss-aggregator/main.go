package main

import (
	"fmt"
	"log"

	"github.com/db-cooper7/bootdev-guided-projects/go-rss-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("config: %+v\n", cfg)

	userName := "Franz Hermann"
	cfg.SetUser(userName)
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("config: %+v\n", cfg)
}
