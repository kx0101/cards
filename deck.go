package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	cards := make(deck, 0, len(cardSuits)*len(cardValues))

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), fs.FileMode(0666))
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeckFromFile(filename string) deck {
	fileContent, err := os.ReadFile(filename)

	if err != nil {
		log.Print(err.Error())
        log.Print("Creating a new deck from scratch...")
		return newDeck()
	}

	s := strings.Split(string(fileContent), ", ")

	return deck(s)
}
