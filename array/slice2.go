package array

import "fmt"

func Ex_slice2() {

	x := []int{1, 2, 3, 4, 5}

	fmt.Println(x[1])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	y := append(x[1:3], 8, 9)
	fmt.Println(y)

}
