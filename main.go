package main

import "github.com/iamarjun/go-pokedex/internal/pokeapi"

type config struct {
	pokeApiClient           pokeapi.Client
	previousLocationAreaUrl *string
	nextLocationAreaUrl     *string
}

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)
}
