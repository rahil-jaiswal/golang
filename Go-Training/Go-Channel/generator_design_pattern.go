package main

import "fmt"

func generate(x ...int) chan int {
	ch := make(chan int, len(x))
	go func() {
		defer close(ch)
		for _, v := range x {
			//fmt.Println(v)
			ch <- v
		}
	}()

	return ch
}

///

func main() {
	xz := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	xCh := generate(xz...)
	for v := range xCh {
		fmt.Println(v)
	}
}
