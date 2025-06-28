package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	// the bot interface establishes a similarity between the englishBot struct and the spanishBot struct - this similarity is the getGreeting function which can be used to call similar methods on each struct that have different logic but return the same type
	// An interface cannot be used as a receiver type because it is not a concrete type like structs, map etc. You cannot create an instance of an interface type. You only define functions (or attributes?) that instances of concrete types can have with it
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (e englishBot) getGreeting() string {
	return "Hello there"
}

func (s spanishBot) getGreeting() string {
	return "Hola, como estas?"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
