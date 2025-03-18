package map_ex

import "fmt"

func Ex_map1() {
	m := map[string]int{
		"yeachan": 30,
		"aaa":     31,
		"bbb:":    32,
	}

	fmt.Println(m)
	fmt.Println(m["yeachan"])
	v, ok := m["aaa"] // 31, true
	fmt.Println(v, ok)

	m["newKey"] = 33 // add element

	for k, v := range m { // loop
		fmt.Println("key is :", k, ", value is : ", v)
	}

	delete(m, "yeachan")

	for k, v := range m { // loop
		fmt.Println("key is :", k, ", value is : ", v)
	}

}
