package main

import "fmt"

func main() {
	brothers_with_name_of_mothers := [][]string{{"raj", "swati"}, {"rahil", "annapurna"}, {"aryan", "ranjeeta"}}
	for i, v := range brothers_with_name_of_mothers {
		for j, w := range v {
			fmt.Printf("%v, %v \n", w, j)
		}
		fmt.Println(i, v, "\n")
	}

	x := []string{"jeet", "mahesh"}
	y := []string{"aashu", "ravi"}
	z := []string{"ayush", "shailesh"}
	brothers_with_name_of_fathers := [][]string{x, y, z}
	fmt.Println(brothers_with_name_of_fathers)
}
