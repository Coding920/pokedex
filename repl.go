package main

import (
	"bufio"
	"fmt"
	"github.com/Coding920/pokedex/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	client *pokeapi.PokeClient
	next   *string
	prev   *string
}

func startRepl(cfg config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		userCmd := cleanedInput[0]

		cmd, ok := getCommands()[userCmd]
		if ok {
			err := cmd.callback(&cfg)
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
