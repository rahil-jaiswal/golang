package main

import "fmt"

func main() {
	//for 1
	for i := 1; i <= 100; i++ {
		if i%2 == 0 && i%3 == 0 && i%5 == 0 {
			fmt.Println(i)
		}
	}

	//for 2
	for {
		fmt.Println("Yeh ek hi baar hoga...")
		break
	}

	//for 3
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j += 3
	}

}
