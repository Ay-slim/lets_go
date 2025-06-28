package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()
	deckSize := len(testDeck)
	if deckSize != 52 {
		t.Errorf("Expected length 52, got %v", deckSize)
	}
	if testDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, got %s", testDeck[0])
	}
	if testDeck[deckSize-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, got %s", testDeck[deckSize-1])
	}
}

func TestWriteToFileAndReadFromFile(t *testing.T) {
	testFilename := "_testingfile"
	os.Remove(testFilename)
	testDeck := newDeck()
	testDeck.writeToFile(testFilename)
	loadedDeck := readFromFile(testFilename)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected file length to be 52, got %v", len(loadedDeck))
	}
	os.Remove(testFilename)
}
