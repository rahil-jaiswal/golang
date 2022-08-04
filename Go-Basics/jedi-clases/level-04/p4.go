package main

import "fmt"

func main() {
	slice := []string{"raj", "rahil", "jeet"}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	//append
	slice = append(slice, "aashu")
	fmt.Println(slice)

	sl2 := []string{"chiku", "ayush"}
	fmt.Println(sl2)

	//append slice
	slice = append(slice, sl2...)
	fmt.Println(slice)

	//delete from slice
	slice = append(slice, "baka")
	fmt.Println(slice)

	slice = append(slice[:len(slice)-1])
	fmt.Println(slice)

}
