package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	openedBuffer, err := os.Open(filename)
	if (err != nil) {
		fmt.Println("Error reading file " + filename )
		return
	}
	io.Copy(os.Stdout, openedBuffer)
}