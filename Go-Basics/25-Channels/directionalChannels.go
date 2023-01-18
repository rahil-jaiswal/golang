package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan int)
	//send
	go foo(c)

	//receive
	go bar(c)
	wg.Wait()
	fmt.Println("Exiting the main")
}

// we can typecast the bidirectional channel into uni-directional channel but not vice-versa
// send
func foo(c chan<- int) {
	c <- 11
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}
