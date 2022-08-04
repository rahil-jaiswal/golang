package main

import (
	"fmt"
)

// Iota is a useful concept for creating incrementing constants in Go.
// iota can be used as const counter
// can be reset with new const block
// starts with 0

const (
	a = iota
	b
	c
)
const (
	d = iota
	e
	f
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("")
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	//bit shifting
	fmt.Println("")
	fmt.Println("")
	x := 4
	fmt.Printf("%v \t\t %b\n", x, x)
	y := x << 1
	fmt.Printf("%v \t\t %b\n", y, y)

	// kb mb gb conversion using iota
	fmt.Println("")
	fmt.Printf("kb %d \t %b\n", kb, kb)
	fmt.Printf("mb %d \t %b\n", mb, mb)
	fmt.Printf("gb %d \t %b\n", gb, gb)
}
