package main

import (
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("provide the name of the pokemon")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("you haven't caught %v yet", pokemon)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Height: %v\n", pokemon.Height)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s, %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Type:")
	for _, typ := range pokemon.Types {
		fmt.Printf("\t- %s\n", typ.Type.Name)
	}

	return nil
}
