package main

import "fmt"

func main() {
	year := 1973
	for {
		if year >= 2001 {
			break
		}
		fmt.Println(year)
		year++
	}
}
