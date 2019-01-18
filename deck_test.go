package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected length to be 52,but found to be %v", len(cards))
	}
	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but found to be %v", cards[0])
	}
	if cards[len(cards)-1] != "King of Clubs" {
		t.Errorf("Expected first card to be King of Clubs but found to be %v", cards[len(cards)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	cards := newDeckFromFile("_decktesting")
	if len(cards) != 52 {
		t.Errorf("Expected length to be 52,but found to be %v", len(cards))
	}
	os.Remove("_decktesting")
}
