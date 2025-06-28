package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main () {
	resp, err := http.Get("http://www.google.com")
	if (err != nil) {
		fmt.Println("error: "+ err.Error())
		os.Exit(1)
	}
	lw := logWriter{}
	// holder := make([]byte, 999999)
	// resp.Body.Read(holder)
	// fmt.Println(string(holder))
	//io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return 1, nil
}