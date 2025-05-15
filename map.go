package main

import (
	"fmt"
)

func commandMapF(cfg *config) error {
	locationsJson, err := cfg.client.GetLocations(cfg.next)
	if err != nil {
		return err
	}

	for _, location := range locationsJson.Results {
		fmt.Println(location.Name)
	}
	cfg.next = locationsJson.Next
	cfg.prev = locationsJson.Previous
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prev == nil {
		fmt.Println("You're on the first page!")
		return nil
	}
	locationsJson, err := cfg.client.GetLocations(cfg.prev)
	if err != nil {
		return err
	}

	for _, location := range locationsJson.Results {
		fmt.Println(location.Name)
	}
	cfg.next = locationsJson.Next
	cfg.prev = locationsJson.Previous
	return nil
}
