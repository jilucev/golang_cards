package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be Ace of Hearts, but was %v", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected first card to be King of Spades, but was %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 25 cards in deck, got %v", len(loadedDeck))
	}
}
