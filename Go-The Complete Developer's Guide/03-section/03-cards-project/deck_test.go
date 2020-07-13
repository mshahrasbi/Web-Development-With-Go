package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// create a new deck
	// write if statement to see if the deck has the right number of cards
	// if it doesn't, tell the go test handler (t) that something went wrong

	d := newDeck()
	if len(d) != 52 {
		t.Errorf("[TestNewDeck] Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("[TestNewDeck] Expected first card 'Ace of Spades' , but got '%v'", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("[TestNewDeck] Expected last card 'King of Clubs' , but got '%v'", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Delete any files in current working directory withthe name "_decktesting"
	// create a deck
	// save to file "_decktesting"
	// load from file
	// assert deck length
	// delete any files in current working directory with the name "_decktesting"

	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("[TestSaveToDeckAndNewDeckFromFile] Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
