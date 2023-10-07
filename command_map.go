package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	url := cfg.nextLocationsURL
	if len(url) == 0 {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	dto, err := cfg.pokeapiClient.GetLocationAreas(url)
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
