package lv3

import "fmt"

func Ex5() {
	x := 10
	for x <= 100 {
		if x%4 == 0 {
			fmt.Println("x can be divided by 4", x)
		}
		x++
	}
}
