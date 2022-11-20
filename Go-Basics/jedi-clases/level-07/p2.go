package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) changeMe() {
	p.first = p.first + "ji"
}
func main() {
	p := person{
		first: "Rahil",
		last:  "Jaiswal",
		age:   25,
	}
	fmt.Println(p)
	p.changeMe()
	fmt.Println(p)

}
