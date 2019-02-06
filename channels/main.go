package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.lk",
		"http://golang.org",
		"http://stackoverflow.com",
		
		"http://amazon.com",
		"http://facebook.com",
	}
	c := make(chan string)
	for _, l := range links {
		go checkLink(l, c)
	}
	//for l:=range c{	
	//	go checkLink(l,c)
	//}
	for l:=range c{	
		go func(link string){
			time.Sleep(3*time.Second)
			checkLink(link,c)
		}(l)
	}
}
func checkLink(link string, c chan string) {
	
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " up")
		c <- link
		return
	}
	fmt.Println(link, " down")
	c <- link
}
