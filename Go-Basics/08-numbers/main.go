package main

import (
	"fmt"
	"runtime"
)

// numeric types
// unsigned integer - uint8 unit16 unit32 unit64
// signed integer - int8 int16 int32 int64
// float - float32 float64
// byte alias of uint8 and rune alias of int32
// uint or int could be by default of 32 or 64 bit depending on compiler if not mentioned explicitly like uint8 or int16

var x int
var y float64

func main() {
	x = 42
	y = 3.124567889942
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	// printing system architecture details
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
