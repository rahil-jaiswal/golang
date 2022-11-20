package main

import "fmt"

type Person struct {
	name string
	age  rune
}

func (p *Person) speaks() {
	fmt.Println("Hi, I am ", p.name, " and I am", p.age)
}

type Human interface {
	speaks()
}

func saySomething(h Human) {
	h.speaks()
}

func main() {
	p1 := Person{"Rahil Jaiswal", 25}
	// does not work
	// saySomething(p1)
	saySomething(&p1)
}
