package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	//receive
	Receive(even, odd, quit)

	fmt.Println("exiting from main")
}

func send(even, odd, quit chan<- int) {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
	close(even)
	close(odd)
	close(quit)
}

func Receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("Even \t", v)
		case v := <-odd:
			fmt.Println("Odd \t", v)
		case v := <-quit:
			fmt.Println("Quit \t", v)
			return
		}
	}
}
