package main

import "fmt"

var x int
var y string
var z bool

func main() {
	// fmt.Printf("%T - %v\n", x, x)
	// fmt.Printf("%T - %v\n", y, y)
	// fmt.Printf("%T - %v\n", z, z)

	s := fmt.Sprintf("%v %v %v ", x, y, z)
	fmt.Println(s)
}
