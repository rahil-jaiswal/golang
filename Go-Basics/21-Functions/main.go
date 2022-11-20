package main

import "fmt"

func main() {
	foo()
	bar("chinu")
	s1 := "Akshita"
	s1 = woo(s1)
	fmt.Println(s1)
	a, b := test("Rahil", "Jaiswal")
	fmt.Println(a)
	fmt.Println(b)
	y := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(y)

	z := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(z...))

	// Anonymously Running....
	func() {
		fmt.Println("Anonymously Running ")
	}()
}

// Everything in go is pass by value and not pass by reference
//function general structure in go
//func (r receiver) funcName(parameters) (returns(s)) { //func body}
// receiver is an attachment to the function of type

func foo() {
	fmt.Println("Hello from foo....")
}

func bar(s string) {
	fmt.Printf("Hey %v, bar welcomes you...\n", s)
}

func woo(s string) string {
	s = s + " Jaiswal"
	return fmt.Sprint(s)
}

func test(first, last string) (string, bool) {
	return fmt.Sprint(first + " " + last), true
}

//variadic parameter
func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
