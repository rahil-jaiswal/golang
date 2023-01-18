package main

import "fmt"

func checkPallindrome(s string) bool {
	var alpha [26]int
	for _, i := range s {
		//fmt.Printf("%c", i)
		alpha[i-rune('a')]++
	}
	var odd rune
	for i := 0; i < 26; i++ {
		if alpha[i]%2 == 1 {
			odd++
		}
	}
	if odd > 1 {
		return false
	}
	return true
}

func main() {
	x := checkPallindrome("abccbaa")
	fmt.Println(x)
}
