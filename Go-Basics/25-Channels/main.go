package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	// bidirectional channels
	c := make(chan int, 3)
	c <- 44
	go func() {
		c <- 42
		fmt.Println(<-c)
		wg.Done()
	}()
	c <- 43
	fmt.Println(<-c)
	wg.Wait()
	fmt.Println(<-c)

	//directional channels
	//read only channels
	i := make(chan<- int, 3)
	//write only channels
	o := make(<-chan int, 3)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", o)
}
