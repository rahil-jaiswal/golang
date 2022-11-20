package main

import (
	"context"
	"fmt"
	"runtime"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel() // cancel when all the work is done

	fmt.Println("GoRoutines :", runtime.NumGoroutine())
	for v := range gen(ctx) {
		fmt.Println(v)
		if v == 100 {
			break
		}
	}

	fmt.Println("GoRoutines :", runtime.NumGoroutine())
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	fmt.Println("GoRoutines :", runtime.NumGoroutine())
	return dst
}
