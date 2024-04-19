// Reducing req time from 350ms to 250ms by using goroutines for fetching user likes and matches concurrently and aggregating the data using channels.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()
	respch := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go fetchUserLikes(userName, respch, wg)
	go fetchUserMatch(userName, respch, wg)
	wg.Wait()
	close(respch)

	for resp := range respch {
		fmt.Println("Response:", resp)
	}

	fmt.Println("User:", userName)
	fmt.Println("Took:", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "KUSH"
}

func fetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	respch <- 11
	wg.Done()
}

func fetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respch <- "ANNA"
	wg.Done()
}
