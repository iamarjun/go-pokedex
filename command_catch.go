package main

import (
	"fmt"
	"log"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("provide the name of the pokemon")
	}

	pokemon := args[0]

	resp, err := cfg.pokeApiClient.GetPokemon(&pokemon)
	
	if err != nil {
		log.Fatal(err)
	}
	
	const threshold = 50
	randNum := rand.Intn(resp.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemon)
	}

	fmt.Printf("%s caught\n", pokemon)

	return nil
}
