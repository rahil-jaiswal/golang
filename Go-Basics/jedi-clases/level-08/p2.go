package main

import (
	"encoding/json"
	"fmt"
)

type Users []struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	jsonRaw := `[{"First":"Raino","Age":25},{"First":"Chini","Age":23},{"First":"Aksu","Age":24}]`

	fmt.Println(jsonRaw)
	var users Users
	err := json.Unmarshal([]byte(jsonRaw), &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
}
