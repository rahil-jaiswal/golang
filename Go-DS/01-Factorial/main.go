package main

import "fmt"

//method 1
func factorial(x int) int {
	if x == 1 {
		return 1
	}
	return x * factorial(x-1)
}

//method 2
func facto(x int) int {
	result := 1
	for i := 1; i <= x; i++ {
		result *= i
	}
	return result
}

func main() {
	a := factorial(6)
	fmt.Println("Method 1 Factorial : ", a)
	fmt.Println("-------------------------")

	b := facto(6)
	fmt.Println("Method 2 Factorial : ", b)
}
