package main

import "fmt"

func main() {
	x, y := foo(bar, "bhad me jao")
	fmt.Printf("%v\n%T\n", x(), x)
	fmt.Printf("%v %T", y, y)
}

func foo(z func() string, s string) (func() string, bool) {

	return func() string {
		return fmt.Sprint("Return a function...." + z() + s)
	}, true
}

func bar() string {
	return fmt.Sprint("Rahil Jaiswal says - ")
}
