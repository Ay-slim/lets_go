package main

import "fmt"

func main() {
	// range lets you iterate over structures like slices, maps, or strings
	test_slice := [] string{"me", "you", "us"}
	for idx, val := range test_slice {
		fmt.Println(idx, val)
	}
	test_map := map [string] int {"a": 1, "b": 2, "c": 3}
	for key, val := range test_map {
		fmt.Println(key, val)
	}
	test_string := "Bah bah black sheep"
	for idx, val := range test_string {
		fmt.Println(idx, val)
	}
}