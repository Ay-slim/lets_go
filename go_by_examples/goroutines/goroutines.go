package main

import "fmt"
import "time"

func sync() {
	for i := 0; i < 3; i++ {
		fmt.Println("Sync", ":", i)
	}
}

func first() {
	for i := 0; i < 3; i++ {
		fmt.Println("First", ":", i)
	}
}

func second() {
	fmt.Println("Second")
}

func main() {
	go second()
	
	go first()

	time.Sleep(time.Second)

	fmt.Println("Done!")
}
