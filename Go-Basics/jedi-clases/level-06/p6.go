package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Anonymously Running...")
	}
	f()

	func() {
		fmt.Println("I am batman...")
	}()
}
