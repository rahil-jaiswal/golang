package main

import "fmt"

func main() {
	foo()        //1
	bar()        //2
	defer cafe() //9
	cafe()       //3
	foo()        //4
	defer foo()  //8
	bar()        //5
	defer cafe() //7
	foo()        //6
}

func foo() {
	fmt.Println("Foooooooooooo")
}

func bar() {
	fmt.Println("Baaaaaaar")
}

func cafe() {
	fmt.Println("caaaaaffffffeeeeee")
}
