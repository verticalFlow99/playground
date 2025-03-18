package function

import "fmt"

func ExFunctionType() {

	var operator func(int, int) int
	operator = getOperator("*")

	result := operator(4, 5)
	fmt.Println("result is", result)
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func add(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}
