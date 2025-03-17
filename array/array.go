package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x[0])
	x[2] = 100
	fmt.Println(x[2])
	fmt.Println(len(x))
}
