package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sitesToCheck := []string{
		"http://www.google.com",
		"http://www.golang.org",
		"http://www.amazon.com",
		"http://somebullshit.com",
	}

	customChannel := make(chan string)

	for _, url := range sitesToCheck {
		go checkSite(url, customChannel)
	}

	// for i := 0; i < len(sitesToCheck); i++ {
	// 	fmt.Println(<-customChannel)
	// }
	// for i := 0; i < len(sitesToCheck); i++ {
	// 	fmt.Println(<-customChannel)
	// }
	for l := range customChannel {
		func(link string) {
			time.Sleep(5 * time.Second)
			go checkSite(link, customChannel)
		}(l)
		go checkSite(l, customChannel)
	}
}

func checkSite(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Site down! print URL:" + url)
		c <- url
	} else {
		fmt.Println("Site up! print URL:" + url)
		c <- url
	}
}
