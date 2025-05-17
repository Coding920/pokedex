package main

import (
	"github.com/Coding920/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := config{
		client:  &pokeClient,
		pokemon: map[string]pokeapi.PokemonData{},
	}
	startRepl(cfg)
}
