package main

import "fmt"

func main () {
	intsArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for intVal := range(intsArr) {
		if intVal % 2 == 0 {
			fmt.Println(intVal, " Even")
		} else {
			fmt.Println(intVal, " Odd")
		}
	}
}