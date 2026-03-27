package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"map": {
			name:			"map",
			description:	"Lists Pokemon location areas",
			callback:		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Lists previous page of location areas",
			callback:		commandMapb,
		},
		"explore": {
			name:			"explore",
			description:	"Lists Pokemon found at location",
			callback:		commandExplore,
		},
		"catch": {
			name:			"catch",
			description:	"Throw a pokeball at a pokemon",
			callback:		commandCatch,
		},
		"inspect": {
			name:			"inspect",
			description:	"Lists given Pokemon's stats",
			callback:		commandInspect,
		},
		"pokedex": {
			name:			"pokedex",
			description:	"Lists the names of Pokemon you have caught",
			callback: 		commandPokedex,
		},
	}


}
func startRepl(cfg *Config) {
	fmt.Print("Pokedex >")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		cmd, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		cmd.callback(cfg, args...)

		}
	}
