package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Printf("Slicing first half  %v\n", x[6:])
	fmt.Printf("Slicing second half %v\n", x[:6])
	fmt.Printf("Slicing in the middle %v\n", x[2:7])
}
