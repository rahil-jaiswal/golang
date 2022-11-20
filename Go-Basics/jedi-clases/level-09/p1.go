package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo(s string) {
	fmt.Println("foo - ", s)
	wg.Done()
}

func bar(s string) {
	fmt.Println("bar - ", s)
	time.Sleep(time.Minute)
	wg.Done()
}

func main() {
	wg.Add(2)
	go foo("second last")
	//defer bar("last")
	go bar("first")
	wg.Wait()

}
