package main // there's something here about not being able to have more than one main function per directory, flagging to revisit after learning about codebase structuring and directory patterns

import "fmt"

func main() {
	// fmt.Println prints each new output on a new line. fmt.Print bunches doesn't separate with new lines
	fmt.Println("hello world")

	// var keyword is used to initialize (must specify type) a variable and/or assign values to it
	var duhh int
	fmt.Println(duhh)
	var mahn int = 3

	// := operator is a shorthand for both initializing and assigning values to a variable (can only be used within a function)
	mohn := "a"
	fmt.Println(mahn, mohn)

  // Of course as the name implies, varibles defined by var or the := operator
	// are mutable and can be reassigned to another value of the same type
	mahn = 8
	mohn = "c"
	fmt.Println(mahn, mohn)

	// Const is used to declare variables whose values are immutable
	const mah string = "p"
	fmt.Println(mah)
}