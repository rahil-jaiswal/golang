package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge ([]Person)

func (b ByAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByAge) Len() int           { return len(b) }
func (b ByAge) Less(i, j int) bool { return b[i].Age > b[j].Age }

func main() {
	p1 := Person{"Rahil", 25}
	p2 := Person{"Raj", 28}
	p3 := Person{"Jeet", 24}
	p4 := Person{"Aashu", 21}
	p5 := Person{"Aryan", 18}
	p6 := Person{"Ayush", 13}

	p := []Person{p1, p2, p3, p4, p5, p6}
	fmt.Println(p)
	sort.Sort(ByAge(p))
	fmt.Println(p)

}
