package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	unsorted := []int{4, 7, 1, 2, 8, 3, 6, 9, 0}
	fmt.Println(unsorted)
	sort.Ints(unsorted)
	fmt.Println(unsorted)

	fmt.Println("-------------")

	unsortedStrings := []string{"Raj", "Rahil", "Jeet", "Aashu", "Aryan", "Ayush"}
	fmt.Println(unsortedStrings)
	sort.Strings(unsortedStrings)
	fmt.Println(unsortedStrings)

	p1 := Person{"Rahil", 25}
	p2 := Person{"Raj", 28}
	p3 := Person{"Jeet", 24}
	p4 := Person{"Aashu", 21}
	p5 := Person{"Aryan", 18}
	p6 := Person{"Ayush", 13}

	p := []Person{p1, p2, p3, p4, p5, p6}
	fmt.Println(p)
	ba := ByAge(p)
	sort.Sort(ba)
	fmt.Println(p)
}
