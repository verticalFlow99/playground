package main

import "fmt"

type mytype int

var x mytype
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 30
	fmt.Println("x is", x)
	y = int(x) // conversion
	fmt.Println("y is", y)
}
