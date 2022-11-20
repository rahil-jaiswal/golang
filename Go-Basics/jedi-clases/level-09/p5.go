package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32
	const gc = 100
	var wg sync.WaitGroup
	wg.Add(gc)

	for i := 0; i < gc; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			fmt.Println(atomic.LoadInt32(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("counter :", counter)
}
