package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = ", MySum(2, 3))
}

func MySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
