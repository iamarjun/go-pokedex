package main

import (
	"time"

	"github.com/iamarjun/go-pokedex/internal/pokeapi"
)

type config struct {
	pokeApiClient           pokeapi.Client
	previousLocationAreaUrl *string
	nextLocationAreaUrl     *string
	caughtPokemon           map[string]pokeapi.PokemonResp
}

func main() {
	cacheInterval := time.Hour
	cfg := config{
		pokeApiClient: pokeapi.NewClient(cacheInterval),
		caughtPokemon: make(map[string]pokeapi.PokemonResp),
	}
	startRepl(&cfg)
}
