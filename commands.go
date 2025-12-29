package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func commandPokedex(cfg *Config, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("This command doesn't take any arguments.")
	}

	fmt.Printf("Your Pokedex: \n")
	for _, each := range cfg.pokedex {
		fmt.Printf("  - %s\n", each.Name)
	}

	return nil
}

func commandInspect(cfg *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Inspect command usage: inspect <pokemon_name>")
	}
	name := args[0]

	info, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("You have not caught that Pokemon yet")
		return nil
	}

	fmt.Printf("Name: %s\n", info.Name)
	fmt.Printf("Height: %d\n", info.Height)
	fmt.Printf("Weight: %d\n", info.Weight)
	fmt.Printf("Stats:\n")
	for _, each := range info.Stats {
		fmt.Printf("    - %s: %d\n", each.Stat.Name, each.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, each := range info.Types {
		fmt.Printf("    - %s\n", each.Type.Name)
	}

	return nil
}

// TODO: make this pretty cause it's a mess
func commandCatch(cfg *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Explore command usage: explore <area_name>")
	}

	res, err := cfg.client.PokemonInfo(args[0])
	if err != nil {
		return fmt.Errorf("error getting pokemon info: %w", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	// fmt.Printf("%s's base experience is: %d\n", args[0], res.BaseExperience)
	time.Sleep(1 * time.Second) // it can't be instant...
	chance := int(80 / (float64(res.BaseExperience) / 100.0))
	// fmt.Println("[DEBUG]: chance is ", chance)
	if chance >= 100 { // instant catch
		fmt.Printf("%s was caught!\n", args[0])
		info, err := cfg.client.PokemonInfo(args[0])
		if err != nil {
			return err
		}
		cfg.pokedex[args[0]] = info
	} else { // otherwise gamble
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		gamba := random.Intn(100)
		if gamba <= chance {
			fmt.Printf("%s was caught!\n", args[0])
			info, err := cfg.client.PokemonInfo(args[0])
			if err != nil {
				return err
			}
			cfg.pokedex[args[0]] = info
		} else {
			fmt.Printf("%s escaped!\n", args[0])
		}

	}

	return nil
}

func commandExplore(cfg *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Explore command usage: explore <area_name>")
	}

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
