package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create new type of 'deck'
// it is a slice of strings
// A slice is like a resizeable array that allows only one data type as elements

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d is basically self/this
// it's called the reciever
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//  d is the parameter and deck is its type
// handSize is the parameter and int is its type
// I am returning two values. Both are of type 'deck'
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// convert instance of deck to string
	// join that into a single, comma-separated string.
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// if the file doesn't already exist, write file will attempt to create it
	// the last argument defines permissions to create that file
	// 0666 means that anyone can do that.
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
