package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	// Test the length of the deck
	expectedDeckSize := 20 // 4 suits * 5 card values
	if len(cards) != expectedDeckSize {
		t.Errorf("Expected deck length of %d, but got %d", expectedDeckSize, len(cards))
	}

	// Test that the first card is "Ace of Spades"
	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", cards[0])
	}

	// Test that the last card is "four of Clubs"
	if cards[len(cards)-1] != "Five of Clubs" {
		t.Errorf("Expected last card to be 'Five of Clubs', but got %v", cards[len(cards)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename) // Cleanup any existing file before starting the test

	deck := newDeck()
	err := deck.saveToFile(filename)
	if err != nil {
		t.Errorf("Failed to save the deck to file: %v", err)
	}

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards in loaded deck, got %v", len(loadedDeck))
	}

	os.Remove(filename) // Cleanup after the test
}

func TestDeal(t *testing.T) {
	deck := newDeck()
	handSize := 5
	hand, remainingDeck := deal(deck, handSize)

	if len(hand) != handSize {
		t.Errorf("Expected hand size of %d, but got %d", handSize, len(hand))
	}

	if len(remainingDeck) != (len(deck) - handSize) {
		t.Errorf("Expected remaining deck size of %d, but got %d", len(deck)-handSize, len(remainingDeck))
	}
}
