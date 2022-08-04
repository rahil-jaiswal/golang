package main

import (
	"fmt"
)

//give error as short declaration variable cannot be used fro global variables
// y := "Global Variable"

func main() {
	x := 10
	fmt.Println(x)
	x = 11
	fmt.Println(x)
	//x = "November" will give error as operand used with short declaration operator cannot change its type once declared
	z := x + 10
	fmt.Println(z)
}
