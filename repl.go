package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

type Config struct {
	Next     *string
	Previous *string
}


var commands map[string]cliCommand

func init() {
    commands = map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
		"map":{
			name: "map",
			description: "It displays the names of 20 location areas in the Pokemon world. Each subsequent call to map will display the next 20 locations, and so on.",
			callback: commandMap,
		},
    }
}


func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	words := strings.Fields(loweredText)
	return words
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: \n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *Config) error {
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	
	location := Locations{}
	err = json.Unmarshal(bytes, &location)
	if err != nil {
    	return err
	}
	
	cfg.Next = location.Next
	cfg.Previous = location.Previous
}
