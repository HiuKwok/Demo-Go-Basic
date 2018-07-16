package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}


	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades %v", d[0] )
	}
}


func TestSaveToDeckAndNewDeckFromFile (t *testing.T){

	fname := "_decktesting"
	os.Remove(fname)

	d := newDeck()
	d.saveToFile(fname)

	loadedDeck := newDeckFromFile(fname)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(fname)


}
