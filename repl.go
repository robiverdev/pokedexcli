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
		fmt.Print("Enter some text: ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		fmt.Println("You've entered:", cleaned)
	}
}

func cleanInput(str string) []string {
	lowerCase := strings.ToLower(str)
	words := strings.Fields(lowerCase)
	return words
}
