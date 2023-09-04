package main

import (
	"fmt"
	"log"
)

func callbackMapBack(cfg *config) error {
	if cfg.previousLocationAreaUrl == nil {
		return fmt.Errorf("you're on the first page")
	}

	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.previousLocationAreaUrl)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas:")

	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.previousLocationAreaUrl = resp.Previous

	return nil
}
