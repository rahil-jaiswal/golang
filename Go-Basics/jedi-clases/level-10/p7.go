package main

import "fmt"

func main() {

	c := make(chan int)

	for j := 1; j <= 10; j++ {
		go func() {
			for i := 1; i <= 10; i++ {
				c <- i * j
			}
		}()
		// close is only needed for range
		//close(c)
	}

	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}

}
