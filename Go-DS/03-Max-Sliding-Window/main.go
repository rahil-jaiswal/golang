package main

import (
	"fmt"
	"sync"
)

func main() {
	x := []int{1, 2, 3, 3, 4, 6, 5, 4, 7}
	w := 3

	y := maxInSlidingWindow(w, x)
	for _, v := range y {
		fmt.Print(v)
	}
	fmt.Println()

	y = maxInSlidingWindow2(w, x)
	for _, v := range y {
		fmt.Print(v)
	}
	fmt.Println()
}

func maxInSlidingWindow(w int, x []int) []int {
	var result []int
	for i := 0; i <= len(x)-w; i++ {
		max := x[i]
		for j := 1; j < w; j++ {
			if x[i+j] > max {
				max = x[i+j]
			}
		}
		result = append(result, max)
	}

	return result
}

func maxInSlidingWindow2(w int, x []int) []int {
	result := make([]int, len(x)-w)
	var wg sync.WaitGroup
	wg.Add(len(x) - w)
	for i := 0; i < len(x)-w; i++ {
		go func(i int) {
			max := x[i+w]
			for j := i; j <= (i + w); j++ {
				if x[j] > max {
					max = x[j]
				}
			}
			result[i] = max
			wg.Done()
		}(i)
	}
	wg.Wait()
	return result
}
