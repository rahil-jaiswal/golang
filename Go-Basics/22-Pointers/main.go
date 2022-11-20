package main

import "fmt"

func main() {
	var x int
	x = 11
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", &x, &x)
	y := &x
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(*&x)

	*y++
	fmt.Println(x)
	foo(x)
}

func foo(x *int) {
	fmt.Println(*x)
}
