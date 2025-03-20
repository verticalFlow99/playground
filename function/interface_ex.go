package function

import "fmt"

type person2 struct {
	first string
	last  string
}

type secretAgent2 struct {
	person2
	ltk bool
}

type human interface {
	speak()
}

func (s secretAgent2) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (p2 person2) speak() {
	fmt.Println("I am", p2.first, p2.last)
}

func bar2(h human) {
	switch h.(type) { // Type Assertion (타입 단언)
	case person2:
		fmt.Println("person2 type", h.(person2).first)
	case secretAgent2:
		fmt.Println("secretAgent2 type", h.(secretAgent2).first)
	}
}

func ExInterface() {
	sa1 := secretAgent2{
		person2: person2{
			first: "yeachan",
			last:  "oh",
		},
		ltk: true,
	}

	p1 := person2{
		first: "Hello",
		last:  "go",
	}

	sa1.speak()
	p1.speak()

	bar2(sa1)
	bar2(p1)

}
