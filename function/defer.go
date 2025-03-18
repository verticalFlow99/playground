package function

import "fmt"

func Ex_defer() {

	defer foo()
	bar()
}

func foo() { fmt.Println("foo") }
func bar() { fmt.Println("bar") }
