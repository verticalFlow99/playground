package main

import "fmt"

func main() {
	x := 1994
	now := 2025
	for x <= now {
		fmt.Println(x)
		x++
	}
}
