package main

import (
	"os"
	"testing"
)

// t - test handler
func TestNewDeck(t *testing.T) {
	testDeck := newDeck()
	if len(testDeck) != 4 {
		t.Errorf("Expected deck length of 4, but got %v", len(testDeck))
	}

	if testDeck[0] != "Bonecrusher Giant" {
		t.Errorf("Expected first card of Bonecrusher Giant, but got %v", testDeck[0])
	}

	if testDeck[len(testDeck)-1] != "Goldspan Dragon" {
		t.Errorf("Expected first card of Goldspan Dragon, but got %v", len(testDeck)-1)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Remove temp test files
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 4 {
		t.Errorf("Expected loaded deck length of 4, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
