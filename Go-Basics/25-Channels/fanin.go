package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go Send(even, odd)

	go receive(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println("Fan In:", v)
	}

	fmt.Println("Exiting from main....")
}

func Send(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func receive(e, o <-chan int, fIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fIn <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fIn <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fIn)
}
