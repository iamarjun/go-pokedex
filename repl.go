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

		fmt.Println(text)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	worlds := strings.Fields(lowered)
	return worlds
}
