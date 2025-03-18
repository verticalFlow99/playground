package function

import "fmt"

func Unfurling() {

	xi := []int{1, 2, 3, 4, 5}
	// sum := sum() // ok
	sum := sum(xi...)
	fmt.Println("sum is : ", sum)

}

func sum(x ...int) int { // unfurling
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
