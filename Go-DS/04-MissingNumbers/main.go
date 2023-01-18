package main

import (
	"fmt"
	"sort"
)

// l. Find missing num in range of numbers.
// var s = []int{1, 2, 3, 4, 6, 7, 8, 9, 10}

func main() {
	var s = []int{13, 12, 11, 15}
	//var s = []int{1, 3, 7, 9, 8, 2, 6, 5}
	findMissingNumber1(s)

}

func findMissingNumber1(s []int) {
	sort.Ints(s)
	min := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != min+i {
			fmt.Println("Missing Number :", min+i)
			return
		}
	}
}
