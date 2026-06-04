package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		for _, cleaned := range cleaned {
			switch cleaned {
			case "exit":
				commandExit()
			case "help":
				commandHelp()
			default:
				fmt.Println("Unknown command")
			}
		}
	}
}