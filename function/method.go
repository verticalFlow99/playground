package function

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func Ex_method() {
	p1 := secretAgent{
		person: person{
			first: "yeachan",
			last:  "oh",
		},
		ltk: true,
	}

	p1.speak()

}
