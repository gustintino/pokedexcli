package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	next     string
	previous string
}

var commands map[string]cliCommand
var cfg config

func initCommands() {
	cfg = config{
		next:     "https://pokeapi.co/api/v2/location-area?offset=0",
		previous: "https://pokeapi.co/api/v2/location-area?offset=0",
	}

	commands = map[string]cliCommand{
		"exit": {
			name:        "Exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "Help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "Map",
			description: "Displays the name of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "Map back",
			description: "Displays the name of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}
