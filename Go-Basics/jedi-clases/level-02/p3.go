package main

import (
	"fmt"
)

const (
	a         = 42
	b float32 = 3.14
)

func main() {
	fmt.Println(a, b)
	fmt.Printf("%T %T", a, b)
}
