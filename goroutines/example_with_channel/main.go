package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func checkUrl(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		// fmt.Println(err)
		// fmt.Printf("%s is DOWN!\n", url)
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		// c <- s // sending into the channel
		fmt.Println(s)
		c <- url

	} else {
		// defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d \n", url, resp.StatusCode)

		// if resp.StatusCode == 200 {
		// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
		// 	file := strings.Split(url, "//")[1]
		// 	file += ".txt"

		// 	s += fmt.Sprintf("Writing response body to %s\n", file)

		// 	err = ioutil.WriteFile(file, bodyBytes, 0664)

		// 	if err != nil {
		// 		// log.Fatal(err)
		// 		s += "Error writing file \n"
		// 		c <- s
		// 	}
		s += fmt.Sprintf("%s is UP\n", url)
		c <- s
		// }
	}
}

func main() {
	pl := fmt.Println

	urls := []string{"https://golang.org", "https://www.google.com", "https://medium.com"}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
		pl(strings.Repeat("#", 20))
	}

	fmt.Println("No. of goroutines: ", runtime.NumGoroutine())

	// for i := 0; i < len(urls); i++ {
	// 	pl(<-c)
	// }
	// for { // infinite loop
	// 	go checkUrl(<-c, c)
	// 	pl(strings.Repeat("#", 20))
	// 	time.Sleep(time.Second * 2)
	// }

	for url := range c {
		time.Sleep(time.Second * 2)
		go checkUrl(url, c)
	}
}
