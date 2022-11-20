package main

import "fmt"

func main() {
	var x string
	fmt.Printf("%v %T\n", x, x)
	x = "Rahil Rocks"
	fmt.Printf("%v %T\n", &x, &x)
}
