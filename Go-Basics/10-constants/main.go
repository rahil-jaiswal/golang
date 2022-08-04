package main

import "fmt"

// individual constants declarations
// const a = 10
// const b = "James Bond"
// const c = 3.14

// constant in single go

const (
	a int8    = 11
	b string  = "Rahil Jaiswal"
	c float64 = 3.142
)

// Un-Typed Constant - no type like int bool float present
// Typed Constat - Value Type like int bool float present

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%v  %T\n", a, a)
	fmt.Printf("%v  %T\n", b, b)
	fmt.Printf("%v  %T\n", c, c)
	const d = 13
	fmt.Println(d)
	// error on below line "cannot assign to d (untyped int constant 13)" as constants cannot be updated
	//d = 12
	//fmt.Println(d)

}
