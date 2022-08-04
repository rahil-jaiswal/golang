package main

import "fmt"

func main() {
	brothers := make([]string, 5, 5)
	fmt.Println(brothers, len(brothers), cap(brothers))
	orig := []string{"raj", "rahil", "jeet", "aashu", "chiku"}
	for i, v := range orig {
		brothers[i] = v
	}

	fmt.Println(brothers, len(brothers), cap(brothers))
	brothers = append(brothers, "ayush")
	fmt.Println(brothers, len(brothers), cap(brothers))
}
