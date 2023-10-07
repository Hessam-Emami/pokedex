package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var args []string
		for i, word := range words {
			if i > 0 {
				args = append(args, word)
			}
		}
		command, exists := GetCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "The map command displays the names of next 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "The map command displays the names of previous 20 location areas in the Pokemon world.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore an area",
			description: "The map command displays the names of the pokemons in the given area.",
			callback:    commandExplore,
		},
	}
}
