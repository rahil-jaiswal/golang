package main

import "fmt"

func main() {
	x, y := foo()
	fmt.Printf("%v\n%T\n", x(), x)
	fmt.Printf("%v %T", y, y)
}

func foo() (func() string, bool) {
	return func() string {
		return fmt.Sprint("Return a function....")
	}, true
}
