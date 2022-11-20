package main

import "fmt"

func main() {

	//anonymous function
	c := make(chan int)

	go func() {
		c <- 11
	}()
	fmt.Println(<-c)

	// buffered channel
	d := make(chan string, 1)
	d <- "rahil"
	fmt.Println(<-d)
}
