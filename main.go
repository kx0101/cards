package main

func main() {
	cards := newDeck()

	cards.shuffle().print()
}
