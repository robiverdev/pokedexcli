package main

import "github.com/robiverdev/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextLocatonAreaURL *string
	previousLocationAreURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
