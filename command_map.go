package main

import (
	"fmt"
	"log"

	"github.com/robiverdev/pokedexcli/internal/pokeapi"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp,err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Result{
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}
