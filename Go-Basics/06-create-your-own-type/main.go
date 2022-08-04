package main

import (
	"fmt"
)

// typical variable declation
var a int
var c int

// creating your own type
type rj int

var b rj

// when creating your own type using type keyword - when converting it into primitive type one need to cast is using
// funtions like int() float() char() - this is called conversions in go

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T \n", a)

	b = 43
	fmt.Println(b)
	fmt.Printf("%T \n", b)

	// this is allowed
	c = a
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	// this is not allowed as b and c are of different types
	// c = b
	// fmt.Println(b)
	// fmt.Printf("%T\n", c)

	// but above can be done using casting
	c = int(b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
