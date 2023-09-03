package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]

		switch command {
		case "help":
			fmt.Println("Welcome to the Pokedex help menu!")
			fmt.Println("Available commands")
			fmt.Println("- help")
			fmt.Println("- exit")
		case "exit":
			os.Exit(0)

		default:
			fmt.Println("Invalid command")
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	worlds := strings.Fields(lowered)
	return worlds
}
