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
	
	_, ok := cfg.caughtPokemon[pokemon]

	if ok {
		return fmt.Errorf("%s already caught", pokemon)
	}

	resp, err := cfg.pokeApiClient.GetPokemon(&pokemon)
	
	if err != nil {
		log.Fatal(err)
	}
	
	const threshold = 50
	randNum := rand.Intn(resp.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemon)
	}

	cfg.caughtPokemon[pokemon] = resp 

	fmt.Printf("%s caught\n", pokemon)

	return nil
}
