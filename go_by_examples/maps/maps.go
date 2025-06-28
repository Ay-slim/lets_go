package main

import (
	"fmt"
	"maps"
)

func main() {
	map_1 := map[string]int{"a": 1, "b": 2}
	map_2 := map[string]int{"a": 1, "b": 2} //map_1
	fmt.Println(map_1, map_2)
	if maps.Equal(map_2, map_1) {
		fmt.Println("Is equal 1")
	}
	map_1["a"] = 3
	if maps.Equal(map_2, map_1) {
		fmt.Println("Is equal 2")
	}
	fmt.Println(map_1, map_2)
}