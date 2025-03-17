package main

import "fmt"

func main() {
	x := 1
	for x <= 10000 {
		fmt.Println(x)
		x++
	}
}
