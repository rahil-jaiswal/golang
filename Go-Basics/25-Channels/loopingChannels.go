package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 1; i < 100; i++ {
			c <- i
		}
		// if you dont close, program will run endlessly
		close(c)
	}()

	//receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
