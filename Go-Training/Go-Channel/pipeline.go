package main

import "fmt"

func Generate(x ...int) chan int {
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

func add(dCh chan struct{}, iCh chan int) chan int {
	aRes := make(chan int)

	go func() {
		defer close(aRes)

		for data := range iCh {
			result := data + 1
			select {
			case <-dCh:
				return
			case aRes <- result:

			}
		}
	}()

	return aRes
}

func multiply(dCh chan struct{}, iCh chan int) chan int {
	mRes := make(chan int)

	go func() {
		defer close(mRes)

		for data := range iCh {
			result := data * 2
			select {
			case <-dCh:
				return
			case mRes <- result:

			}
		}
	}()

	return mRes
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	doneCh := make(chan struct{})

	defer close(doneCh)
	ch := Generate(x...)
	resCh := multiply(doneCh, add(doneCh, ch))

	for v := range resCh {
		fmt.Println(v)
	}
}
