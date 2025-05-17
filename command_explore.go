package main

import (
	"fmt"
)

func explore(cfg *config, params []string) error {
	if len(params) < 2 {
		return fmt.Errorf("Didn't provide location to explore")
	}

	locationInfo, err := cfg.client.ExploreLocation(params[1])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v\n", params[1])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationInfo.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
