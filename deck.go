package main

import (
	"fmt"
)

// create a new type of deck which is slice string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four",
		"Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack",
		"Queen", "King"}

	for _, suits := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suits)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
