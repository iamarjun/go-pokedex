package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	for k, _ := range cfg.caughtPokemon {
		fmt.Println(k)
	}
	return nil
}
