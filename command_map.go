package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocatonAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Result {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocatonAreaURL = resp.Next
	cfg.previousLocationAreURL = resp.Previous
	return nil
}

func callbackMapPrev(cfg *config) error {
	if cfg.previousLocationAreURL == nil {
		return errors.New("You're on the first page")
	}
	resp,err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Result{
		fmt.Printf(" - %s\n",area.Name)
	}
	cfg.nextLocatonAreaURL = resp.Next
	cfg.previousLocationAreURL = resp.Previous
	return nil
}

