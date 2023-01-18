package main

import (
	"fmt"
	"sync"
)

func main() {

	// scenario 3
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	defer close(ch)
	wg.Add(1)
	go func() {
		ch <- 4
		ch <- 5
		ch <- 6
		wg.Done()
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// if we leave unbuffered channel with elements it will give error but same works for buffer channel
	// fmt.Println(<-ch)

	//// scenario 2
	//// receiving from empty buffered channel creates deadlock
	//wg.Add(1)
	//ch2 := make(chan int, 2)
	//defer close(ch2)
	//go func() {
	//	fmt.Println(<-ch2)
	//	wg.Done()
	//}()

	// scenario 3
	// sending to empty buffered channel and not receiving gives deadlock
	//wg.Add(1)
	//ch3 := make(chan int, 2)
	//defer close(ch3)
	//go func() {
	//	ch <- 11
	//	wg.Done()
	//}()

	//// scenario 4
	//// only receiving from buffered channel and not sending gives deadlock
	//wg.Add(1)
	//ch4 := make(chan int, 4)
	////ch4 <- 1
	//defer close(ch4)
	//go func() {
	//	fmt.Println(<-ch4)
	//	wg.Done()
	//}()

	// scenario 5
	// sending to empty unbuffered channel and not receiving gives deadlock
	//wg.Add(1)
	//ch3 := make(chan int)
	//defer close(ch3)
	//go func() {
	//	ch <- 11
	//	wg.Done()
	//}()

	//// scenario 6
	//// only receiving from unbuffered channel and not sending gives deadlock
	//wg.Add(1)
	//ch4 := make(chan int)
	////ch4 <- 1
	//defer close(ch4)
	//go func() {
	//	fmt.Println(<-ch4)
	//	wg.Done()
	//}()

	//// scenario 7
	//// sending more elements than the size of buffered channel gives error
	//fmt.Println("Scenario 7\n")
	//wg.Add(1)
	//ch5 := make(chan int, 3)
	//go func() {
	//	ch5 <- 5
	//	ch5 <- 6
	//	ch5 <- 7
	//	ch5 <- 8
	//	wg.Done()
	//}()
	//fmt.Println(<-ch5)
	//fmt.Println(<-ch5)

	wg.Wait()
}
