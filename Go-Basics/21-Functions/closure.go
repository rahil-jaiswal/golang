package main

import "fmt"

// closure - closing values into a parent function for a returning function

func parent() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	a := parent()
	b := parent()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(a())

}
