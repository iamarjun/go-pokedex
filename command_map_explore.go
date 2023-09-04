package main

import (
	"fmt"
	"log"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("not location area provided")
	}

	locationAreaName := args[0]
	resp, err := cfg.pokeApiClient.GetLocationArea(locationAreaName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pokemons in areas:")

	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
