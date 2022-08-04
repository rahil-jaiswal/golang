package main

import "fmt"

//maps are key value store
//super fast super efficient
//unordered list

func main() {
	x := map[string]int{
		"Raj":   27,
		"Rahil": 25,
		"Jeet":  24,
	}

	fmt.Println(x)
	fmt.Println(x["Rahil"])

	//trying to extract key which is not present - it will always return zero
	fmt.Println(x["Chiku"])
	v, ok := x["aashu"]
	fmt.Println(v, ok)
	v, ok = x["Jeet"]
	fmt.Println(v, ok)

	//adding element to a map
	x["Aashish"] = 20
	fmt.Println(x)

	if v, ok := x["Aashish"]; ok {
		fmt.Println("Aashish hai...   ", v)
	}

	//looping on maps
	for k, v := range x {
		fmt.Println(k, v)
	}

	fmt.Println("")
	fmt.Println(x)
	//delete from a mao
	delete(x, "Rahil")
	fmt.Println(x)
}
