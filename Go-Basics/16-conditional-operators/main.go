package main

import "fmt"

func main() {
	fmt.Printf("true && true \t\t %v\n", true && true)
	fmt.Printf("false && true \t\t %v\n", false && true)
	fmt.Printf("true && false \t\t %v\n", true && false)
	fmt.Printf("true && false \t\t %v\n", false && false)
	fmt.Printf("true || true \t\t %v\n", true || true)
	fmt.Printf("false || true \t\t %v\n", false || true)
	fmt.Printf("true || false \t\t %v\n", true || false)
	fmt.Printf("true && false \t\t %v\n", false || false)
	fmt.Printf("!false \t\t %v\n", !false)
	fmt.Printf("!true \t\t %v\n", !true)
}
