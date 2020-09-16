package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// a slice of strings
type deck []string

// Create and return a list of deck
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
// works like this in JS
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Return a deck list of string
func newDeckFromFile(filename string) deck {
	//byteslice and error
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		// exit the program
		os.Exit(1)
	}

	// convert the byte slice to a slice of string
	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {

	generateRandomNumber := time.Now().UnixNano()

	source := rand.NewSource(generateRandomNumber)
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
