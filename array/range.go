package array

import "fmt"

func Ex_range() {

	x := []int{1, 2, 3, 4, 5}
	for i, v := range x { // use range for loop
		fmt.Println(i, v)
	}
}
