package main

import "fmt"

func main() {
	cs := make(chan int)
	go func(chan<- int) {
		cs <- 10
		close(cs)
	}(cs)

	fmt.Println(<-cs)

	fmt.Printf("Channel Type %T\n", cs)
}
