package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)

	fmt.Println("Hand:")
	hand.print()

	fmt.Println("Remaining Deck:")
	remainingDeck.print()
}

