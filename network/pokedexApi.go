package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationAreaResponseDto struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []LocationAreaDto `json:"results"`
}

type LocationAreaDto struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PageState struct {
	Next     string
	Previous string
}

// why should it be *Client?
func (c Client) GetLocationAreas(uri string) (LocationAreaResponseDto, error) {
	var locationAreaResponseDto LocationAreaResponseDto
	resCache, isAvailable := c.cache.Get(uri)
	if isAvailable {
		err := json.Unmarshal(resCache, &locationAreaResponseDto)
		if err != nil {
			log.Fatalf("Parsing exception from the Cache %s", err.Error())
		} else {
			return locationAreaResponseDto, nil
		}
	}
	res, err := http.Get(uri)
	if err != nil {
		return LocationAreaResponseDto{}, errors.New("Network problem. Check the network and try again")
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return LocationAreaResponseDto{}, errors.New("Unknown issue. Check the network and try again")
	}
	if err != nil {
		log.Fatal(err)
		return LocationAreaResponseDto{}, errors.New("Unknown issue. Check the network and try again")
	}
	c.cache.Add(uri, body)
	err = json.Unmarshal(body, &locationAreaResponseDto)
	if err != nil {
		log.Fatalf("Parsing exception %s", err.Error())
		return LocationAreaResponseDto{}, errors.New("Network problem. Check the network and try again")
	}
	return locationAreaResponseDto, nil
}

func (c Client) ExploreArea(area string) (LocationAreaDetailDto, error) {
	var locationAreaDetailDto LocationAreaDetailDto
	uri := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)
	resCache, isAvailable := c.cache.Get(uri)
	if isAvailable {
		err := json.Unmarshal(resCache, &locationAreaDetailDto)
		if err != nil {
			log.Fatalf("Parsing exception from the Cache %s", err.Error())
		} else {
			return locationAreaDetailDto, nil
		}
	}
	res, err := http.Get(uri)
	if err != nil {
		return LocationAreaDetailDto{}, errors.New("Network problem. Check the network and try again")
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return LocationAreaDetailDto{}, errors.New("Unknown issue. Check the network and try again")
	}
	if err != nil {
		log.Fatal(err)
		return LocationAreaDetailDto{}, errors.New("Unknown issue. Check the network and try again")
	}
	c.cache.Add(uri, body)
	err = json.Unmarshal(body, &locationAreaDetailDto)
	if err != nil {
		log.Fatalf("Parsing exception %s", err.Error())
		return LocationAreaDetailDto{}, errors.New("Network problem. Check the network and try again")
	}
	return locationAreaDetailDto, nil
}

type LocationAreaDetailDto struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
