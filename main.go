package main

import (
	"time"

	"github.com/gustintino/pokedexcli/internal/pokeapi"
)

type Config struct {
	client *pokeapi.Client

	next     *string
	previous *string
}

func main() {
	baseLink := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	cfg := Config{
		client:   pokeapi.NewClient(5 * time.Second),
		next:     &baseLink,
		previous: nil,
	}
	getCommands()
	startRepl(cfg)
}
