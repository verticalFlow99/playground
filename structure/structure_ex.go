package structure

import "fmt"

type person2 struct {
	first          string
	last           string
	favoriteColors []string
}

func Example() {

	practice1()
	practice3()
	practice4()
}

func practice1() {

	fmt.Println("--- practice 1 ---")

	p1 := person2{
		first: "firstname1",
		last:  "lastname1",
		favoriteColors: []string{
			"white", "black",
		},
	}

	p2 := person2{
		first: "firstname2",
		last:  "lastname2",
		favoriteColors: []string{
			"orange", "green",
		},
	}

	fmt.Printf("%+v\n", p1)
	fmt.Printf("%+v\n", p2)

	practice2(p1, p2)
}

func practice2(p1 person2, p2 person2) {

	fmt.Println("--- practice 2 ---")

	m := map[string]person2{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, color := range v.favoriteColors {
			fmt.Println("index :", i, color)
		}
	}
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func practice3() {
	fmt.Println("--- practice 3 ---")
	v1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: false,
	}

	v2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1.doors)
	fmt.Println(v2.doors)

}

func practice4() {

	fmt.Println("--- practice 4 ---")

	myAnonymousStruct := struct {
		field1 int
		field2 string
		field3 bool
	}{
		field1: 1,
		field2: "hello",
		field3: true,
	}

	fmt.Println("myAnonymousStruct is: ", myAnonymousStruct)

}
