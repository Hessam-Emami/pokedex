package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *config, args ...string) error {
	if len(cfg.prevLocationsURL) == 0 {
		return errors.New("You are at the first page. proceed forward!")
	}
	dto, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = dto.Next
	cfg.prevLocationsURL = dto.Previous
	for _, area := range dto.Results {
		fmt.Println(area.Name)
	}
	fmt.Println()
	return nil
}
