package main

import "fmt"

func main() {
	x := [5]int{}
	for i := len(x) - 1; i >= 0; i-- {
		x[i] = i
	}
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y := [6]int{27, 25, 24, 20, 18, 12}
	for i, v := range y {
		fmt.Println(y[i], v, i)
	}

	fmt.Printf("%T\n", y)
}
