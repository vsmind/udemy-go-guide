package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkSite(url, c)
	}

	/*for i := 0; i < len(urls); i++ {
		fmt.Println(<- c)
	}*/

	fmt.Println("**********Start**********")
	//Infinite loop
	//for  {
	//	go checkSite(<- c, c)
	//}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkSite(link, c)
		}(l)
	}
}

func checkSite(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "might be down.")
		c <- url
		return
	}

	fmt.Println(url, "is up and running.")
	c <- url
}
