package main

import "fmt"

type person struct {
	first string
	last  string
	age   rune
}

func (p person) speak() {
	fmt.Println("He is ", p.first, p.last, "of age", p.age, "and he speaks too...")
}

func main() {
	p1 := person{
		first: "Rahil",
		last:  "Jaiswal",
		age:   25,
	}

	//fmt.Println(p1)
	p1.speak()

}
