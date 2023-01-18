package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			fmt.Println("i:", x)
			wg.Done()
		}(i)
	}

	y := "Darth Vader"
	fmt.Printf("y : %v & %v\n", &y, y)
	wg.Add(1)
	go func() {
		fmt.Printf("y : %v & %v\n", &y, y)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main ends...")
}
