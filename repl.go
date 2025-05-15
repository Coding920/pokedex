package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		command := cleanedInput[0]
		fmt.Println("Your command was: " + command)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
