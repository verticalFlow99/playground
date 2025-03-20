package function

import "fmt"

func ExCallback() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	evenNumber := even(addElements, slice...)
	oddNumber := odd(addElements, slice...)

	fmt.Println("add all even numbers :", evenNumber)
	fmt.Println("add all odd numbers :", oddNumber)

}

func addElements(s ...int) int {
	sum := 0
	for i, _ := range s {
		sum += s[i]
	}

	return sum
}

func even(f func(s ...int) int, xi ...int) int {
	var yi []int
	for _, v := range xi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}

func odd(f func(s ...int) int, xi ...int) int {
	var yi []int
	for _, v := range xi {
		if v%2 == 1 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}
