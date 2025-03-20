package function

import "fmt"

type Stringer interface { // 1. Stringer 인터페이스 선언
	String() string
}

type Student struct {
	name string
	age  int
}

type Teacher struct {
	subject string
	age     string
}

func (s Student) String() string { // 2. Student 타입에 String() 메서드 추가
	return fmt.Sprintf("hello I'm %d years old and name is %s", s.age, s.name)
}

func (t Teacher) String() string { // 2. Student 타입에 String() 메서드 추가
	return fmt.Sprintf("hello I'm %s years old and my major is %s", t.age, t.subject)
}

func jobChecker(s Stringer) {
	switch s.(type) { // 타입 단언
	case Student:
		fmt.Println("I'm student")
	case Teacher:
		fmt.Println("I'm teacher")
	}
}

func ExInterface2() {

	student := Student{
		name: "yeachan",
		age:  30,
	}
	var stringer Stringer
	stringer = student            // 3. 변수 stringer 의 값으로 student 대입
	fmt.Println(student.String()) // 4. 변수 stringer 의 String() 메서드 호출

	teacher := Teacher{
		subject: "english",
		age:     "55",
	}

	var stringer2 Stringer
	stringer2 = teacher
	fmt.Println(teacher.String())

	jobChecker(stringer)
	jobChecker(stringer2)

}
