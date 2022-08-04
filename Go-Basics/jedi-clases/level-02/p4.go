package main

import "fmt"

func main() {
	x := 3
	fmt.Printf("%d - %b - %#X\n", x, x, x)
	y := x << 1
	fmt.Printf("%d - %b - %#X\n", y, y, y)

}
