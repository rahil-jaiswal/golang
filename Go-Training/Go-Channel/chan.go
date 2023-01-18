package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	eCh := make(chan int)
	oCh := make(chan int)
	go func(eCh, oCh chan int) {
		for i := 1; i <= 20; i++ {
			if i%2 == 0 {
				eCh <- i
			} else {
				oCh <- i
			}
		}
		close(eCh)
		close(oCh)
	}(eCh, oCh)

	go func() {
		receive(eCh)
		wg.Done()
	}()
	go func() {
		receive(oCh)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main done...")
}

func receive(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func receive2(even, odd <-chan int) {
	for {
		select {
		case v := <-odd:
			fmt.Println("Odd :", v)
		case v := <-even:
			fmt.Println("Even : ", v)
		}

	}
}
