package main

// array - fixed size
// slice - dynamic size
// all elements in an array are of the same type
// all elements in a slice are of the same type
// array is a value type
// slice is a reference type
// array is a fixed size
// slice is a dynamic size
// array is a value type
// slice is a reference type
func main() {
	cards := deck{newCard(), "ace of spades"}
	cards = append(cards, "six of diamonds")

	// cards.print()
	createdDeck := newDeck()
	// newHand, remainingDeck := deal(createdDeck, 5)
	// newHand.print()
	// remainingDeck.print()
	// fmt.Println(createdDeck.toString())
	// createdDeck.writeToFile("cards_file")
	createdDeck.print()
	// createdDeck.shuffle()
	// createdDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
