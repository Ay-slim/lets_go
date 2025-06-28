package main

import "fmt"

func ival (val int) {
	val = 0
}

func iptr (ptr *int) {
	*ptr = 0
}

func main() {
	i := 1
	fmt.Println(i, " should print 1")

	// value change (futile)
	ival(i)
	fmt.Println(i, " should print 1")

	// value change (fruitful)
	iptr(&i)
	fmt.Println(i, " should print 0")
}
