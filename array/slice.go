package main

import "fmt"

func main() {

	// x := []type{values}
	arr1 := []int{1, 2, 3, 7, 8}
	fmt.Println(len(arr1))

	arr2 := arr1[1:2] // slice
	fmt.Println(arr1)
	fmt.Println(arr2)

	arr1 = append(arr1, 3, 4, 5) // add elements to array
	fmt.Println(arr1)

	arr3 := make([]int, 3, 5)
	fmt.Println(arr3)
	copy(arr1, arr2)
	fmt.Println(arr2)

	//for i, v := range x {
	//	fmt.Println("i is :", i, " v is : ", v)
	//}
}
