package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{6, 3, 7, 4, 1, 8, 2, 9, 0}
	fmt.Println("Unsorted =>", x)
	sort.Ints(x)
	fmt.Println("Sorted   =>", x)

	fmt.Println("-----------------")

	y := []string{"a", "z", "h", "k", "n", "v", "i", "l", "r"}
	fmt.Println("Unsorted =>", y)
	sort.Strings(y)
	fmt.Println("Sorted   =>", y)

}
