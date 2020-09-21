package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://phrecommends.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// receive message from channel
	// every time your run a go routine, need to map how many routines you are expecting to be executed
	for l := range c {
		// fire an anonymous funcion
		go func(l string) {
			// pause the go routing for 5 sec
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, " might be down!")
		// send message to channel
		c <- l

		return
	}
	fmt.Println(l, " is up!")
	// send message to channel
	c <- l
}
