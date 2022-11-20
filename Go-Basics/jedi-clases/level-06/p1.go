package main

import "fmt"

func main() {
	a := foo()
	b, c, d := bar()
	fmt.Println(a)
	fmt.Println(b, c, d)

}

func foo() int {
	return 11
}

func bar() (int, string, bool) {
	return 11, "November", true
}
