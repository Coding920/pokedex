package main

import (
	"fmt"
)

func inspect(cfg *config, params []string) error {
	if len(params) < 2 {
		return fmt.Errorf("Which pokemon do you want to inspect?")
	}

	pokemon, ok := cfg.pokemon[params[1]]
	if !ok {
		return fmt.Errorf("you haven't caught %v yet", params[1])
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, innerType := range pokemon.Types {
		fmt.Printf("  - %v\n", innerType.InnerType.Name)
	}

	return nil
}
