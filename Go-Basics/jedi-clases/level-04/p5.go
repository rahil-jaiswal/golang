package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)

	//delete from middle of slice
	x = append(x[:4], x[5:]...)
	fmt.Println(x)

	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	//deleting first element of slice
	x = append(x[1:])
	fmt.Println(x)

	//deleting last element of slice
	x = append(x[:len(x)-1])
	fmt.Println(x)
}
