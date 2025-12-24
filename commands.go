package main

import (
	"fmt"
	"github.com/gustintino/pokedexcli/internal/pokeapi"
	"os"
)

func commandMap(cfg *config) error {
	res, err := pokeapi.GetLocationAreas(cfg.next)
	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}

	for _, each := range res.Results {
		fmt.Println(each.Name)
	}

	cfg.next = res.Next
	cfg.previous = res.Previous

	return nil
}

func commandMapb(cfg *config) error {
	pokeapi.GetLocationAreas(cfg.previous)
	res, err := pokeapi.GetLocationAreas(cfg.previous)
	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}

	for _, each := range res.Results {
		fmt.Println(each.Name)
	}

	cfg.next = res.Next
	cfg.previous = res.Previous
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
