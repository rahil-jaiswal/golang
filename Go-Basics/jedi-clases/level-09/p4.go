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

	var mu sync.Mutex

	for i := 0; i < gc; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("counter :", counter)
}
