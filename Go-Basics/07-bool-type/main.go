package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	a := 12
	b := 12
	if a == b {
		fmt.Println("a equals b")
	}

}
