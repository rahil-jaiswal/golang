package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Go OS\t\t", runtime.GOOS)
	fmt.Println("Go ARCH\t\t", runtime.GOARCH)
	fmt.Println("Go Processors\t\t", runtime.NumCPU())
	fmt.Println("Go Routines\t\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	//wg.Wait()
	bar()
	fmt.Println("Go Processors\t\t", runtime.NumCPU())
	fmt.Println("Go Routines\t\t", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 11; i++ {
		fmt.Println("foo ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 11; i++ {
		fmt.Println("bar ", i)
	}
	wg.Wait()
	//wg.Done()
}
