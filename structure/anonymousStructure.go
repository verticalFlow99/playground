package structure

import "fmt"

func Ex_anonymousStructure() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "yeachan",
		last:  "oh",
		age:   30,
	}

	fmt.Println(p1)
}
