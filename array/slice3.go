package array

func Ex_slice3() {

	x := make([]int, 10, 11)
	for i := 0; i < 3; i++ {
		x = append(x, i)
		println(x)
	}
	//	[11/11]0x14000108ed8
	//	[12/22]0x1400012e000
	//	[13/22]0x1400012e000
	//fmt.Println(x)
	//fmt.Println("len:", len(x)) // 현재 저장된 요소의 개수
	//fmt.Println("cap:", cap(x)) // 저장할 수 있는 최대 개수

	//for i := range x {
	//	x[i] = i + 1
	//	println(x)
	//}

	//fmt.Println(x)
}
