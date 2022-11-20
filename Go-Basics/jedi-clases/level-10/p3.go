package main

import "fmt"

func main() {
	c := gen()
	receive(c)

	fmt.Println("Done & Dusted")
}

func gen() <-chan int {
	x := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			x <- i
		}
		close(x)
	}()

	return x
}

func receive(y <-chan int) {
	fmt.Println("Receiving channel input")
	for v := range y {
		fmt.Println(v)
	}
}
