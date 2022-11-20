package main

import "fmt"

func main() {
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(facto(5))
}

func facto(x int) int {
	if x == 0 {
		return 1
	}
	return x * facto(x-1)
}
