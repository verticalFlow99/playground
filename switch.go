package main

import "fmt"

func main() {
	x := 5
	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("this is not ture, but will be run")
	default:
		fmt.Println("default")
	}

}
