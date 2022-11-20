package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

type rectangle struct {
	length  float32
	breadth float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float32 {
	return r.length * r.breadth
}

type shape interface {
	// return type is important
	area() float32
}

func Info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c1 := circle{
		radius: 11.234,
	}

	r1 := rectangle{
		length:  4.56,
		breadth: 7.65,
	}

	// fmt.Println(r1.area())
	// fmt.Println(c1.area())

	Info(c1)
	Info(r1)

}
