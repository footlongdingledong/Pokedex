package main

import (
	"strings"
	"fmt"
	"os"
	"github.com/footlongdingledong/pokedexcli/internal/pokeapi"
	"math/rand"
)


var commands map[string]cliCommand
var cfg *Config


func cleanInput(text string) []string {
	split := strings.Fields(strings.ToLower(text))
	return split
}

func commandExit(cfg *Config, args...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args ...string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for name, cmd := range commands {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}

func mapList(location pokeapi.LocationArea) {
	for i:= 0; i < len(location.Results); i++ {
		fmt.Println(location.Results[i].Name)
	}
}

func commandMap(cfg *Config, args...string) error {
	location, err := cfg.pokeapiClient.GetLocations(*cfg.Next)
	if err != nil {
		return err
	}
	mapList(location)
	cfg.Next = location.Next
	cfg.Previous = location.Previous
	return nil
}


func commandMapb(cfg *Config, args ...string) error {
	if cfg.Previous == nil {
		fmt.Print("you're on the first page\n")
		return nil
	}
	location, err := cfg.pokeapiClient.GetLocations(*cfg.Previous)
	if err != nil {
		return err
	}
	mapList(location)
	cfg.Next = location.Next
	cfg.Previous = location.Previous
	return nil
}

func commandExplore(cfg *Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Command requires name of location")
		return nil
	}
	local := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + local
	location, err := cfg.pokeapiClient.GetLocationInfo(url)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for i := range location.PokemonEncounters {
		fmt.Println(" -"+location.PokemonEncounters[i].Pokemon.Name)
	}
	return nil
}


func commandCatch(cfg *Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Command requires name of Pokemon")
		return nil
	}
	pokemon := args[0]
	pokeinfo, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil{
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	chance := (325 - pokeinfo.BaseExperience)/(325/100)
	roll := rand.Intn(100) <= chance
	if roll {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.caughtPokemon[pokemon] = pokeinfo
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}
	return nil
}

func commandInspect(cfg *Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Command requires name of Pokemon")
		return nil
	}
	pokemon := args[0]
	mon, ok := cfg.caughtPokemon[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", mon.Name, mon.Height, mon.Weight)
		for i := 0; i < len(mon.Stats); i++ {
			fmt.Printf("  -%s: %d\n", mon.Stats[i].Stat.Name, mon.Stats[i].BaseStat)
		}
		fmt.Println("Types:")
		for i := 0; i < len(mon.Types); i++ {
			fmt.Println("  -" + mon.Types[i].Type.Name)
		}
	}
	return nil
}

func commandPokedex(cfg *Config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You do not have any pokemon")
	}
	fmt.Println("Your Pokedex:")
	for i := range cfg.caughtPokemon {
		fmt.Println(" - " + cfg.caughtPokemon[i].Name)
	}
	return nil
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*Config, ...string) error
}

type Config struct {
	pokeapiClient	pokeapi.Client
	caughtPokemon	map[string]pokeapi.Pokemon
	Next			*string
	Previous 		*string
}
