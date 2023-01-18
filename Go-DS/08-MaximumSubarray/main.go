package main

import "fmt"

func maxSumSubArray(a []int) []int {
	// check if length is one then return same array
	if len(a) == 1 {
		return a
	}
	// set current and max sum as 0
	maxSum := 0
	currSum := 0
	var begin int
	var end int
	var start int
	for i := 0; i < len(a); i++ {
		// add current element to current sum
		currSum += a[i]
		// check if current sum is less than current element
		// if yes then set new begin at current element and set current sum to 0
		if currSum < a[i] {
			currSum = a[i]
			begin = i
		}

		// check if current sum is greater than max sum
		// if yes update max sum as well as start and end index
		if maxSum < currSum {
			maxSum = currSum
			start = begin
			end = i
		}
	}
	fmt.Println(maxSum)
	return a[start : end+1]
}

func main() {
	fmt.Println(maxSumSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
