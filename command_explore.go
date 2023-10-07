package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Location name is empty!")
	}
	detail, err := cfg.pokeapiClient.ExploreArea(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", detail.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range detail.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
