package main

import (
	"fmt"
)

func pokedex(cfg *config, params []string) error {
	fmt.Println("Your Pokedex:")

	for name, _ := range cfg.pokemon {
		fmt.Printf(" - %v\n", name)
	}

	return nil
}
