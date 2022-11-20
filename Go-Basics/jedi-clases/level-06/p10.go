package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func foo() func() int {
	var x rune
	return func() int {
		x++
		return int(x)
	}
}
