package main

import "fmt"

type myvar int

var x myvar = 21
var y int

func main() {
	fmt.Println(x)
	x = 22
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
