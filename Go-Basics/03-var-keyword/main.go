package main

import "fmt"

// go is static programming language i.e. once the variable is declared its type cannot be changed dynamically
//unlike short declaration operation var can be used for global variables
var x = 20

// these are raw string literals.
// strings can be declared with both "double quotes" and `back quotes`
var a = `Rahil said "Chini, Would you like to have coffee with me?"`
var b = "Chini replied, `Yeah, will catchup soon...`"

func main() {
	fmt.Println(x)
	// assigned another value
	x = 21
	// following will give error as var type is changing to string which is not supported
	// x = "Rahil Jaiwal"
	fmt.Println(x)
	//one using short declaration operator
	y := true
	// declared variable of type but havent used it
	var z bool
	z = false
	fmt.Println(y, z)

	fmt.Println(a)
	fmt.Println(b)

}
