package main

import (
	"fmt"
	"os"
)

func commandExplore(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Explore command usage: explore <area_name>")
	}

	// FIXME: something is out of range

	res, err := cfg.client.AreaDetails(args[0])
	if err != nil {
		return fmt.Errorf("Error getting area details: %w", err)
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Printf("Found Pokemon:\n")
	for _, each := range res.PokemonEncounters {
		fmt.Printf("- %s\n", each.Pokemon.Name)
	}

	return nil
}

func commandMap(cfg *Config, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("The Map command doesn't take any arguments")
	}

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

func commandMapb(cfg *Config, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("The Map Back command doesn't take any arguments")
	}

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

func commandExit(cfg *Config, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("The exit command doesn't take any arguments")
	}

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("The help command doesn't take any arguments")
	}

	fmt.Printf("\n")
	// NOTE: i do not want to call the command again, but i don't see an elegant fix right now so i'll leave it
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
