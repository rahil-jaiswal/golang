package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "Rahil",
		last:  "Jaiswal",
		favFlavors: []string{
			"chocolate",
			"vanilla",
			"butterscotch",
		},
	}

	p2 := person{
		first: "Akshita",
		last:  "Jaiswal",
		favFlavors: []string{
			"butterscotch",
			"chocolate",
		},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	personMap := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}

	for k, v := range personMap {
		fmt.Printf("%v >> %v\n", k, v)
	}

	p2.first = "Chini"
	personMap["Akshita"] = p2
	fmt.Println(personMap["Akshita"])
}
