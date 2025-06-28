package main

import "fmt"

func main() {
	// for loops are the only kind of Go loops
	// for feels intuitive coming from both Javascript and Python background
	i := 1
	fmt.Println("i")
	for (i <= 3) {
		fmt.Println(i)
		i++
	}
	fmt.Println("j")
	for j := 2; j < 6; j++ {
		fmt.Println(j)
	}
	fmt.Println("k")
	for k:=range(4) {
		fmt.Println(k)
	}
}