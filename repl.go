package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		args := []string{}

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(cfg, args...)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	worlds := strings.Fields(lowered)
	return worlds
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "map",
			description: "Lists previous location areas",
			callback:    callbackMapBack,
		},
		"explore": {
			name:        "explore",
			description: "explore an area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon",
			callback:    callbackCatch,
		},
	}
}
