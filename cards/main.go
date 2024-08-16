package main

// import "fmt"

func main() {
	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// cards.saveToFile("my_cards")
	var cards deck = newDeckFromFile("my_cards")
	cards.print()
}

