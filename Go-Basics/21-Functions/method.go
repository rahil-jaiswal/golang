package main

import "fmt"

type contact struct {
	email  string
	mobile int64
}

type person struct {
	name string
	contact
	handicap bool
}

// func (r receiver) identifier(parameters) (returns(s)) { code }

func (p person) speak() {
	fmt.Println("I am ", p.name, ", call me @ ", p.mobile)
}

func main() {
	p1 := person{
		name: "Rahil Jaiswa",
		contact: contact{
			email:  "rj783325@gmail.com",
			mobile: 9403382619,
		},
		handicap: false,
	}

	//fmt.Println(p1)
	p1.speak()
}
