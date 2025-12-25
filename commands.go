package main

import (
	"fmt"
	"os"
)

func commandMap(cfg *Config) error {
	if cfg.next == nil {
		fmt.Println("You are at the end of the list, there are no next pages")
		return nil
	}

	res, err := cfg.client.GetLocationAreas(*cfg.next)
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

func commandMapb(cfg *Config) error {
	if cfg.previous == nil {
		fmt.Println("You are at the beginning of the list, there are no previous pages")
		return nil
	}

	res, err := cfg.client.GetLocationAreas(*cfg.previous)
	if err != nil {
		fmt.Printf("the prev url is: %s", *cfg.previous)
		return fmt.Errorf("error getting location areas: %w", err)
	}

	for _, each := range res.Results {
		fmt.Println(each.Name)
	}

	cfg.next = res.Next
	cfg.previous = res.Previous

	return nil
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Printf("\n")
	// NOTE: i do not want to call the command again, but i don't see an elegant fix right now so i'll leave it
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
