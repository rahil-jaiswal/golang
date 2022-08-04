package main

import "fmt"

func main() {
	// x := type{values} - composite literals
	// slice allows you to group together the values of same type

	x := []int{4, 5, 6, 7, 9, 42}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) //probably capacity
	fmt.Println(x[4])

	for i, v := range x {
		fmt.Printf("{ %d , %d }\n", i, v)
	}

	// slicing the slices using : operator
	y := []int{11, 12, 45, 78, 43, 80}
	fmt.Println(y)
	fmt.Println(y[:])
	fmt.Println(y[1:3])
	fmt.Println(y[:4])
	fmt.Println(y[3:])

	//append to a slice
	x = append(x, 1, 33, 44, 54, 67)
	fmt.Println(x)

	//append one slice to aonther slice.
	x = append(x, y...)
	fmt.Println(x)

	//delete from middle of slice
	x = append(x[:3], x[7:]...)
	fmt.Println(x)

	//delete first element of slice
	x = append(x[:0], x[1:]...)
	fmt.Println(x)

	//delete last element of slice
	x = append(x[:len(x)-1])
	fmt.Println(x)

	// make function - is used by underlying go to create arrays and slices
	// make(type, size, capacity)
	a := make([]int, 5, 6)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 2)
	fmt.Println(a, len(a), cap(a))

	//multi-dimesional slices
	m := []int{1, 2, 3}
	n := []int{4, 5, 6}
	fmt.Println(m)
	fmt.Println(n)
	o := [][]int{m, n}
	fmt.Println(o)
}
