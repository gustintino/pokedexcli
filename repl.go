package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg Config) {
	fmt.Println("Welcome to the Pokedex!")

	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		cleanInput := cleanInput(scanner.Text())
		if len(cleanInput) == 0 {
			fmt.Println("Please type a command")
			continue
		}
		command := cleanInput[0]

		if value, ok := commands[command]; ok {
			err := value.callback(&cfg)
			if err != nil {
				fmt.Println(fmt.Errorf("error occured: %w", err))
			}
		} else {
			fmt.Println("Command wasn't found, please try again.")
		}

	}
}

func cleanInput(text string) []string {
	split := strings.Fields(text)
	var result []string
	for _, each := range split {
		result = append(result, strings.ToLower(each))
	}

	return result
}
