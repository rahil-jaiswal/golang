package main

import "fmt"

// { } are compulsory with if else
func main() {
	//if
	if true {
		fmt.Println("yes")
	}
	if false {
		fmt.Println("no")
	}
	if !true {
		fmt.Println("not yes")
	}
	if !false {
		fmt.Println("not no")
	}

	//if with initilalization statement
	// here x has limited scope of if block
	if x := 42; x >= 42 {
		fmt.Println("one with initialization in if statement")
	}

	//if-else ladder
	if x := 42; x <= 40 {
		fmt.Println("x is smaller ", x)
	} else if x == 42 {
		fmt.Println("x is ok ", x)
	} else {
		fmt.Println("x is bigger ", x)
	}

}
