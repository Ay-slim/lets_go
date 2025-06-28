package main

import "fmt"

type strIntMap map[string] int

func main() {
	sampleMap := strIntMap {
		"a": 1,
		"b": 2,
		"c": 3,
	}
	printMap(sampleMap)
}

func printMap(m strIntMap) {
	for key, val := range m {
		fmt.Println(key, val)
	}
}