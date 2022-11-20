package main

import (
	"fmt"
)

type contact struct {
	email  string
	mobile int64
}

type person struct {
	name string
	contact
	handicap bool
}

type agent struct {
	sign string
	doa  bool
}

// func (r receiver) identifier(parameters) (returns(s)) { code }

func (p person) speak() {
	fmt.Println("I am ", p.name, ", call me @ ", p.mobile)
}

func (a agent) speak() {
	st := "dead"
	if a.doa {
		st = "alive"
	}
	fmt.Println(a.sign, "", st)
}

// interface

type human interface {
	speak()
	//die()
}

func speakUp(h human) {
	h.speak()
	switch h.(type) {
	case person:
		// assertion - typecasting a generic interface into a struct type temporarily (depending on conditions) for compiler to read it as that struct
		fmt.Println(h.(person).name)
	case agent:
		fmt.Println(h.(agent).sign)

	}
}

func main() {
	p1 := person{
		name: "Rahil Jaiswal",
		contact: contact{
			email:  "rj783325@gmail.com",
			mobile: 9403382619,
		},
		handicap: false,
	}

	//fmt.Println(p1)
	speakUp(p1)
	a1 := agent{
		sign: "2p",
		doa:  true,
	}
	speakUp(a1)
}
