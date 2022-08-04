package main

import "fmt"

// zero value is value of a declared but not initialized variable
// for example y here is empty by default
// int x is zero by default before any value is assigned to it
// Int 0
// String ""
// Float/Number 0.0
// nil for Pointer Slice Maps Channels Interfaces Functions
// nil is equivalent to null

var y string
var x int

func main() {
	fmt.Println("printing value of y : ", y, "ending")
	fmt.Printf("%T\n", y)
	y = "Baymax"
	fmt.Println("printing value of y : ", y, "ending")
	fmt.Printf("%T\n", y)

	fmt.Println("printing value of x : ", x, "ending")
	fmt.Printf("%T\n", x)
	x = 11
	fmt.Println("printing value of x : ", x, "ending")
	fmt.Printf("%T\n", x)

}
