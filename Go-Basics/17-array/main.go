package main

import "fmt"

func main() {
	var x [5]int //[5]int is its own type
	var y [6]int //[6]int is its own type and is different from [5] int

	fmt.Println(x)
	fmt.Println(y)
	x[4] = 1
	y[1] = 1

	fmt.Println(x)
	fmt.Println(y)
	// length of array can be determined using len() function
	fmt.Println(len(x), len(y))
}
