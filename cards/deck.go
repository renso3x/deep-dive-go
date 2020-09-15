package main

import "fmt"

// Create a new type of 'deck'
// a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamond", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cs := range cardSuits {
		for _, cv := range cardValues {
			cards = append(cards, cv+" of "+cs)
		}
	}
	return cards
}

// function receiver just 1-3letter for initializing
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
