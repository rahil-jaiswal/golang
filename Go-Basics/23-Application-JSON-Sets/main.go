package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Person []struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Rahil",
		Last:  "Jaiswal",
		Age:   25,
	}

	p2 := person{
		First: "Akshita",
		Last:  "Jaiswal",
		Age:   25,
	}

	p := []person{p1, p2}
	fmt.Println(p)
	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	rawJson := `[{"First":"Rahil","Last":"Jaiswal","Age":25},{"First":"Akshita","Last":"Jaiswal","Age":25}]`
	var p3 Person
	err = json.Unmarshal([]byte(rawJson), &p3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p3)

}
