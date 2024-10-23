package main

import "fmt"

func main() {
	cards := newDeck() // Create a new deck
	cards.shuffle()    // Shuffle the deck
	cards.print()      // Print the shuffled deck

	hand, remainingDeck := deal(cards, 5) // Deal a hand of 5 cards
	hand.print()                          // Print the dealt hand
	remainingDeck.print()                 // Print the remaining deck

	err := cards.saveToFile("my_deck.txt") // Save the deck to a file
	if err != nil {
		fmt.Println("Error saving deck:", err)
	}

	loadedDeck := newDeckFromFile("my_deck.txt") // Load the deck from the file
	loadedDeck.print()                           // Print the loaded deck

}
