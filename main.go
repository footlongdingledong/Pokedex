package main

import(
	"time"
	"github.com/footlongdingledong/pokedexcli/internal/pokeapi"
)


func main() {
	pokeClient := pokeapi.NewClient(10 * time.Second, 5 * time.Minute)
	initialURL := "https://pokeapi.co/api/v2/location-area/"
	cfg = &Config {
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		Next: &initialURL,
		Previous: nil,
	}
	startRepl(cfg)
}
