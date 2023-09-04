package main

import (
	"fmt"
	"log"
)

func callbackMap(cfg *config, args ...string) error {

	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.nextLocationAreaUrl)

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
