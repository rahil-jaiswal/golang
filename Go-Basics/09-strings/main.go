package main

import "fmt"

func main() {
	s := "Hello, Rahil Jaiswal"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	//byte string - slice of byte
	bs := []byte(s)
	fmt.Println(bs)
	// slice of ASCII - UTF-8 encoding for string s
	fmt.Printf("%T\n", bs)

	//Printing UTF -8 values of Characters
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println()

	//Printing Hexadecimal code
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c \t %v \t %#x \n", s[i], s[i], s[i])
	}
}
