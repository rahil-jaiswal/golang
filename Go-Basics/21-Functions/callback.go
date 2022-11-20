package main

import "fmt"

// callback is passing a function as a parameter
// not recommended in go but good to know about

func main() {
	ii := []int{1, 2, 3, 4, 5}
	s := sum(ii...)
	fmt.Println(s)

	iv := even(sum, ii...)
	fmt.Println(iv)
}

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// callback is passing a function as argument to a function parameter
func even(f func(x ...int) int, z ...int) int {
	var result []int
	for _, v := range z {
		if v%2 == 0 {
			result = append(result, v)
		}
	}

	return f(result...)
}
