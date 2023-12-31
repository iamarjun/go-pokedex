package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Available commands")
	fmt.Println("")

	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf("- %s: %s\n", command.name, command.description)
	}

	fmt.Println("")

	return nil
}
