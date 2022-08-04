package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Printf("%d - %T\n", x, x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#X\n", x)
}
