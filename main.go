package main

import (
	"time"

	"github.com/Hessam-Emami/pokedex/network"
)

type config struct {
	pokeapiClient    network.Client
	nextLocationsURL string
	prevLocationsURL string
}

func main() {
	config := &config{
		pokeapiClient: network.NewClient(10*time.Second, 5*time.Minute),
	}
	startRepl(config)
}
