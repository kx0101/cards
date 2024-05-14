package main

func main() {
	cards := newDeck()

	cards.shuffle().print()

    cards.saveToFile("output.txt")
}
