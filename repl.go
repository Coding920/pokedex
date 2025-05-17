package main

import (
	"bufio"
	"fmt"
	"github.com/Coding920/pokedex/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	client  *pokeapi.PokeClient
	next    *string
	prev    *string
	pokemon map[string]Pokemon
}

type Pokemon struct {
	name           string
	baseExperience int
}

func startRepl(cfg config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputWords := cleanInput(scanner.Text())
		userCmd := inputWords[0]

		cmd, ok := getCommands()[userCmd]
		if ok {
			err := cmd.callback(&cfg, inputWords)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
