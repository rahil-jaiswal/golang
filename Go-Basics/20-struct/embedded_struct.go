package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	tataTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "maroon",
		},
		fourWheel: false,
	}

	polo := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(polo)
	fmt.Println(tataTruck)
}
