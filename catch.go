package main

import (
	"fmt"
	"math/rand"
)

func catch(cfg *config, params []string) error {
	if len(params) < 2 {
		return fmt.Errorf("Which pokemon are you trying to catch?")
	}
	pokeData, err := cfg.client.CatchPokemon(params[1])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v\n", params[1])
	if rand.Intn(600) < pokeData.BaseExperience { // From searching, 600 should be a bit over the max of 563
		fmt.Printf("%v escaped\n", params[1])
		return nil
	}

	fmt.Printf("%v was caught!\n", params[1])
	// Todo add pokemon to user's pokedex

	return nil
}
