package main

import "fmt"

func main() {
	// 3rd
	defer foo(" jaiswal")
	bar()
	// 2nd
	defer foo("rahil")
	bar()
	// 1st
	defer foo("chini")
	bar()
	bar()
}

// defer - to put off
// to leave something until later time.
// defer always runs on any exit point of function body
// when there are multiple differ - they are stored in a list and executed in last in first out fashion
func foo(x string) {
	fmt.Println("Foo", x)
}

func bar() {
	fmt.Println("Bar")
}
