package main

import "fmt"

func main() {

	x := foo(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("sum is", x)

	y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	z := foo(y...)
	fmt.Println("sum is", z)
	u := bar(y)
	fmt.Println("sum is", u)

}

// variadic parameter
func foo(x ...int) int {
	total := 0
	for _, i := range x {
		total += i
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
