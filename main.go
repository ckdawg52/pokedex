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
		fmt.Print("Your command was: " + cleaned[0])
	}
}