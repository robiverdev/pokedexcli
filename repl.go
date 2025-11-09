package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter some text: ")
		scanner.Scan()
		text := scanner.Text()

		fmt.Println("You've entered:", text)
	}
}
