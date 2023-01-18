package main

import "fmt"

func square(x int) {
	fmt.Println(x * x)
}

// functions are first class citizens in go
// can be passed as paramteter/argument
// can be returned
// can be assigned to variable
// in all above case its an anonymous function 

func main() {
	// assigning a function to a variable
	f := func() {
		fmt.Println("How are you doing?")
	}

	f()

	g := func(x int32) {
		fmt.Println(x * x)
	}
	g(11)

	h := bar()
	x := h()
	fmt.Println(x)
}

// returning a function from a function
func bar() func() int {
	return func() int {
		return 1997
	}
}
