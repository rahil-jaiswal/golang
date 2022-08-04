package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(v, i, sum)
	}

	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
