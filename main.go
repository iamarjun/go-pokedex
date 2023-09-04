package main

import (
	"time"

	"github.com/iamarjun/go-pokedex/internal/pokeapi"
)

type config struct {
	pokeApiClient           pokeapi.Client
	previousLocationAreaUrl *string
	nextLocationAreaUrl     *string
}

func main() {
	cacheInterval := time.Hour
	cfg := config{
		pokeApiClient: pokeapi.NewClient(cacheInterval),
	}
	startRepl(&cfg)
}
