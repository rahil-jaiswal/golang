package main

import "fmt"

//b. Find count of occurance digit 2 in the range of number ie: n=35
//Eg: 1-10-> 1, 11-20->2, 21-30->10

func main() {
	var x, y, n int
	_, err := fmt.Scanf("%d\n%d", &x, &y)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}
	c := countOfDigit(x, y, n)
	fmt.Println("Count :", c)
}

func countOfDigit(x, y, n int) int {
	if y < x {
		return -1
	}
	var result rune
	for i := x; i <= y; i++ {
		result += countDigitsInNumber(n, i)
	}

	return int(result)
}

func countDigitsInNumber(x, y int) rune {
	var result rune
	for i := y; i > 0; {
		if i%10 == x {
			result++
		}
		i = i / 10
	}
	return result
}
