package main

import "fmt"

type contact struct {
	email  string
	mobile string
}

// type person
type person struct {
	firstname string
	lastname  string
	contact
}

func main() {
	c3 := contact{
		email:  "rj@gmail.com",
		mobile: "11111997",
	}
	// creating a value of type person
	p1 := person{
		firstname: "rahil",
		lastname:  "jaiswal",
		contact:   c3,
	}

	p2 := person{
		firstname: "chini",
		lastname:  "jaiswal",
		contact: contact{
			email:  "chinikam@gmail.com",
			mobile: "11042000",
		},
	}

	fmt.Println(p1, p2)
	fmt.Printf("%T\n", p1)
	c1 := contact{
		email:  "rahiljaiswal@gmail.com",
		mobile: "7875873404",
	}

	fmt.Println(c1)
	fmt.Printf("%T\n", c1)

	p2.firstname = "akshu"
	fmt.Println(p2)

	//anonymous struct
	p3 := struct {
		firstname string
		lastname  string
		age       int32
	}{
		firstname: "sahil",
		lastname:  "kartoli",
		age:       31,
	}

	fmt.Println(p3)

}
