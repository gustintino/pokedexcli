package main

import (
	"time"

	"github.com/gustintino/pokedexcli/internal/pokeapi"
)

// TODO: load the commands into the config here in main?
type Config struct {
	client  *pokeapi.Client
	pokedex map[string]pokeapi.PokemonResponse

	next     *string
	previous *string
}

func main() {
	baseLink := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	cfg := Config{
		client:   pokeapi.NewClient(5 * time.Second),
		pokedex:  make(map[string]pokeapi.PokemonResponse),
		next:     &baseLink,
		previous: nil,
	}
	startRepl(cfg)
}
