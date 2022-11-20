package main

import "fmt"

func main() {
	//creating anonymous struct
	s := struct {
		first     string
		last      string
		nickname  string
		education string
		height    float32
		hobbies   []string
		age       rune
	}{
		first:     "Akshita",
		last:      "Jaiswsl",
		nickname:  "Chini",
		education: "Microbiology",
		height:    5.8,
		hobbies: []string{
			"Dancing",
			"Cooking",
			"Netflix",
			"Reading",
		},
		age: 23,
	}

	fmt.Println(s)
	fmt.Println("We call her ", s.nickname)
}
