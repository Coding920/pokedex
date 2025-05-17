package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, params []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists locations, with each consecutive call to map giving more locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous locations, with each consecutive call to mapb going back further",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Lists Pokemon that can be found at the location provided. Usage: explore (location)",
			callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "Usage: 'catch (pokemon)', has a chance to catch pokemon, based on base experience of the pokemon",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Usage: 'inspect (pokemon)', provides pokemon info for pokemon that you have caught",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists pokemon that you've caught",
			callback:    pokedex,
		},
	}
}

func commandExit(cfg *config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, params []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}

	return nil
}
