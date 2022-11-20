package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{"Raino", 25}
	u2 := user{"Chini", 23}
	u3 := user{"Aksi", 23}
	users := []user{u1, u2, u3}
	fmt.Println(users)
	br, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(br))

}
