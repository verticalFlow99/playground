package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	fmt.Println("Hello World")
	_, _ = fmt.Println("string", 20, true)
	foo()
}

func foo() {
	fmt.Println("hello foo")
}
