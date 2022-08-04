package main

import "fmt"

func main() {
	year := 1997
	for year < 2001 {
		fmt.Println(year)
		year++
		if year >= 2001 {
			break
		}
	}

}
