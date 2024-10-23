
# Go Card Deck

  

This project is a simple implementation of a card deck using Go. It allows you to create a deck of cards, shuffle them, deal a hand, and save/load the deck to/from a file. The deck consists of a combination of four suits (Spades, Diamonds, Hearts, and Clubs) and five card values (Ace, Two, Three, Four, Five).

  

## Features

  

- Create a new deck of cards.

- Print the deck to the console.

- Shuffle the deck randomly.

- Deal a specified number of cards from the deck.

- Convert the deck to a string format.

- Save the deck to a file and load it back from a file.

  

## Getting Started

  

### Prerequisites

  

To run this project, you need to have Go installed on your machine. You can download and install Go from the [official website](https://golang.org/doc/install).

  

### Installation

  

1. Clone the repository:

  

```bash
git clone https://github.com/yourusername/go-card-deck.git

cd go-card-deck
```
2. Run the program:
```go
go run main.go
```
## Usage
Here is a brief description of each function in the program:

-   **Create a New Deck**
    
  To create a new deck, use the `newDeck()` function:
  
```go
cards := newDeck()
```
    
-   **Print the Deck**
    
  To print out each card in the deck along with its index, use the `print()` function:
```go
cards.print()
```
- **Shuffle the Deck**

  To randomize the order of the cards in the deck, call the `shuffle()` method:
 ```go
 cards.shuffle()
 ```
- **Deal a Hand**

  To deal a hand of cards, use the `deal()` function. It returns two decks: the hand and the remaining deck:
```go
 hand, remainingDeck := deal(cards, 5)
```
- **Save the Deck to a File**

  You can save the deck to a file using the `saveToFile()` method:
```go
err := cards.saveToFile("my_deck.txt")
```
  If an error occurs while saving the file, it will return the error.
- **Load a Deck from a File**

  To load a deck from a file, use the `newDeckFromFile()` function:
```go
loadedDeck := newDeckFromFile("my_deck.txt")
```
   This function reads from a specified file and returns a new deck. If the file doesn't exist, it will exit the program.
### Project Structure

The project consists of the following files:
     
		.
		├── deck.go       # Contains all the functions related to the deck operations.
		├── main.go       # The entry point of the application.
		├── deck_test.go  # Unit tests for the deck.
		└── README.md     # Project documentation.
### Example

Here’s an example of how to use the deck functions:
package main
```go
func main() {
    cards := newDeck()          // Create a new deck
    cards.shuffle()             // Shuffle the deck
    cards.print()               // Print the shuffled deck

    hand, remainingDeck := deal(cards, 5)  // Deal a hand of 5 cards
    hand.print()               // Print the dealt hand
    remainingDeck.print()       // Print the remaining deck

    err := cards.saveToFile("my_deck.txt") // Save the deck to a file
    if err != nil {
	fmt.Println("Error saving deck:", err)
    }

    loadedDeck := newDeckFromFile("my_deck.txt") // Load the deck from the file
    loadedDeck.print()            // Print the loaded deck
}
```

		
