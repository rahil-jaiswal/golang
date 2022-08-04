package main

import "fmt"

func main() {
	x := 25
	a := (x == 25)
	b := (x >= 24)
	c := (x <= 24)
	d := (x < 25)
	e := (x > 25)
	f := (x != 0)

	fmt.Println(a, b, c, d, e, f)
}
