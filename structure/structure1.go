package structure

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person // embedded field (내장된 필드)
	ltk    bool
}

func Ex_structure1() {
	sa1 := secretAgent{
		person: person{
			first: "firstname",
			last:  "lastname",
			age:   99,
		},
		ltk: true,
	}

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "yeachan",
		last:  "oh",
		age:   30,
	}

	fmt.Println(sa1.person.first, sa1.person.last, sa1.person.age)
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}
