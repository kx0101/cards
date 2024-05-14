package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())

	cards.saveToFile("output.txt")

	// cards.print()

	// hand, remainingDeck := deal(cards, 5)
	//
	//     hand.print()
	//     remainingDeck.print()
}
