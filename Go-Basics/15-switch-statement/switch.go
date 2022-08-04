package main

import "fmt"

func main() {
	//swtich 1
	switch {
	case false:
		fmt.Println("Nahi Chapega")
	case 2 > 4:
		fmt.Println("Firr bhi nahi chapega")
	case 3 == 3:
		fmt.Println("Chapega Chapega")
	case true:
		fmt.Println("Agar upar wala chapa toh ye nahi chapega")
	}

	//switch 2 with fallthrough
	switch {
	case false:
		fmt.Println("nope")
	case true:
		fmt.Println("yas")
		fallthrough
	case 2 == 4:
		fmt.Println("shayad yes if fall through")
		fallthrough
	case 3 == 2:
		fmt.Println("yaha bhi fallthrough")
		fallthrough
	case 5 != 0:
		fmt.Println("final through")
	case true:
		fmt.Println("but no fallthrough toh nahi chapega")
	}

	//switch 3 with default
	switch {
	case 2 == 3:
		fmt.Println("nahi chapega")
		fallthrough
	case 3 != 4:
		fmt.Println("ye bhi nahi chapega")
		fallthrough
	default:
		fmt.Println("default chapega")
	}

	//switch on var
	x := 10
	y := 3
	switch y {
	case 2, 3:
		fmt.Printf("binary - %b", x)
	case 8:
		fmt.Printf("octal - %o", x)
	case 16:
		fmt.Printf("hexa - %#x", x)
	default:
		fmt.Printf("decimal - %d", x)
	}

	fmt.Println("")
}
