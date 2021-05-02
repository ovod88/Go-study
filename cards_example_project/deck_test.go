package main

import (
	"os"
	"testing"
)

func TestNewdeck(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
