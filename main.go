package main

import (
	"github.com/Coding920/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		client: &pokeClient,
	}
	startRepl(cfg)
}
