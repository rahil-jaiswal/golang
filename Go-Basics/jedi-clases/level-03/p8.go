package main

import "fmt"

func main() {
	favSport := "kabaddi"
	switch favSport {
	case "cricket":
		fmt.Println("Is Dhoni your favorite player?")
	case "basketball":
		fmt.Println("Kobe Bryant isn't it?")
	case "kabaddi":
		fmt.Println("Sandeep Narwal?")
		fallthrough
	case "chess":
		fmt.Println("Vishwanathan Anand")
	default:
		fmt.Println("Ab kounsa khel khel rahe ho?")
	}
}
