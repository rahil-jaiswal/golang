package main

import "fmt"

func main() {
	x := map[string][]string{
		"raj":   {"january", "civil", "07"},
		"rahil": {"november", "computer", "11"},
		"jeet":  {"april", "commerce", "22"},
	}

	for i, v := range x {
		fmt.Printf("%v %v\n", i, v)
	}

	fmt.Println(x)

	//add a record to the map
	x["aashu"] = []string{"august", "bodybuilder", "1"}
	fmt.Println()
	fmt.Println(x)

	//add another record to map
	x["varun"] = []string{"october", "cheating", "17"}
	fmt.Println("")
	fmt.Println(x)

	//deleting a record from map
	delete(x, "varun")
	fmt.Println("")
	fmt.Println(x)

}
