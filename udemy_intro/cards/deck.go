package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for idx, card := range d {
		fmt.Println(idx, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := deck{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := deck{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(deckOfCards deck, handSize int) (deck, deck) {
	return deckOfCards[:handSize], deckOfCards[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) writeToFile(filename string) error {
	bytified := []byte(d.toString())
	return os.WriteFile(filename, bytified, 0666)
}

func readFromFile(filename string) deck {
	fileBytes, error := os.ReadFile(filename)
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}
	splitString := strings.Split(string(fileBytes), ",")
	return deck(splitString)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for idx, _ := range d {
		randIdx := r.Intn(len(d))
		tmpVal := d[idx]
		d[idx] = d[randIdx]
		d[randIdx] = tmpVal
	}
}
