package main

import (
	"fmt"
	"math/rand"
)

func catch(cfg *config, params []string) error {
	if len(params) < 2 {
		return fmt.Errorf("Which pokemon are you trying to catch?")
	}
	pokeData, err := cfg.client.GetPokemon(params[1])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", params[1])
	if rand.Intn(pokeData.BaseExperience+100) < pokeData.BaseExperience {
		fmt.Printf("%v escaped\n", params[1])
		return nil
	}

	fmt.Printf("%v was caught!\n", params[1])
	cfg.pokemon[params[1]] = pokeData

	return nil
}
