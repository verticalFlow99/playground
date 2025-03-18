package fuction

import "fmt"

func Ex_recursiveCall(n int) {
	fmt.Println("현재값:", n)
	if n == 0 {
		return
	}
	Ex_recursiveCall(n - 1)
}
