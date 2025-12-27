package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, args []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"explore": {
			name:        "Explore",
			description: "Lists all of the Pokemon located in the given area. Accepts an <area_name> parameter",
			callback:    commandExplore,
		},
	}
}
