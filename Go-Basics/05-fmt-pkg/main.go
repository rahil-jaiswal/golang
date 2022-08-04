package main

import (
	"fmt"
)

// package fmt implements  formatted I/O  with functions analogues to C' printf and scanf
// Group 1 : Print Printf Println : Printing Output directly to console
// Group 2 : Sprint Sprintf Sprintln : printing to a string which can then assigned to a variable
// Group 3 : Fprint Fprintf : printing to a file or a webserver response or anything that implements the writer interface
var x = 42

func main() {
	// Group 1
	fmt.Printf("Default Format %v\n", x)
	fmt.Printf("Type %T\n", x)
	fmt.Printf("Decimal %d\n", x)
	fmt.Printf("Binary %b\n", x)
	fmt.Printf("Octal %o\n", x)
	fmt.Printf("Hexa %x\n", x)
	fmt.Printf("Hexadecimal %#x\n", x)

	// Group 2
	y := fmt.Sprint(x)
	fmt.Printf("y : %s\n", y)
	fmt.Printf("%T\n", y)

}
