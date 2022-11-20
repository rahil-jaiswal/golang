package main

import "fmt"

var g func()

func main() {
	f := func() {
		fmt.Println("Assigning to a varialbe....")
	}

	f()
	fmt.Printf("Type of variable %T\n", f)
	//fmt.Printf("Value of varialbe %v", f)

	g = f
	g()
}
