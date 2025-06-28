package main

import "fmt"
// Structs, think of these as objects or classes with properties and methods
type base struct {
	num int
}

func (b base) someMethod() int {
	fmt.Println("The number is: %d", b.num)
	return b.num
}

type container struct {
	base
	name string
}

func main() {
	con := container {
		base {num: 3},
		"boo",
	}
	con.base.someMethod()
}
