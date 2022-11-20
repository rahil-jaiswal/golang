package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0
	const gc = 100
	var wg sync.WaitGroup
	wg.Add(gc)
	for i := 0; i < gc; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	fmt.Println("counter :", counter)
	wg.Wait()
}
