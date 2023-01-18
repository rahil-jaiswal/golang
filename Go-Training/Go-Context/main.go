package main

import (
	context "context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	context := context.WithValue(context.Background(), "character", "darth vader")
	wg.Add(1)
	fmt.Println(context.Value("character"))
	fmt.Println(context)

	go task1(context)
	wg.Wait()
}

func task1(ctx context.Context) {
	fmt.Println("Task 1 in progress...")
	timeoutContext, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()
	go task2(timeoutContext)
	fmt.Println("Task 1 back in progress...")
	//time.Sleep(time.Second * 5)
	wg.Done()
	fmt.Println("Task 1 completed...")
}

func task2(ctx context.Context) {
	fmt.Println("Executing Task 2...")
	time.Sleep(time.Second * 20)
	ctx.Done()
	fmt.Println("Task 2 Completed...")
}
