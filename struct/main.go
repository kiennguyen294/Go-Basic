package main

import "fmt"

// struct la kieu du lieu nguoi dung tu dinh nghia
type StudentInfo struct {
	class string
	email string
	age int
}
type Student struct {
	id int
	name string
}

type Student_Nested struct {
	id int
	name string
	info StudentInfo
}
func main()  {
	// khai bao named ==> nen su dung cho tuong minh
	st1 := Student{
		id : 123,
		name: "Ken",
	}
	fmt.Println(st1)
	fmt.Println(st1.id)
	fmt.Println(st1.name)

	// Khai bao ko can name
	st2 := Student{456, "Robin"}
	fmt.Println(st2)
	fmt.Println(st2.id)
	fmt.Println(st2.name)

	//
	var st3 Student = struct {
		id   int
		name string
	}{
		id: 777,
		name: "Bill",
	}

	fmt.Println(st3)
	fmt.Println(st3.id)
	fmt.Println(st3.name)


	// anonymous struct : struct vo danh
	var anonymous = struct {   // khong can khai bao ten cua struct
		email string
		age int
	}{
		"bill@gmail.com",
		27,
	}
	fmt.Println(anonymous)


	// pointer tro toi struct
	pointer := &Student{
		789, "Bill",
	}
	fmt.Println(pointer) //
	fmt.Println(&pointer)
	fmt.Println((*pointer).id)
	fmt.Println((*pointer).name)

	 // anonymous fields
	type NoName struct {
		string
		int
	}
	n := NoName{"Nguyen Van A", 6666}
	fmt.Println(n)

	// struct long struct - Nested struct
	student01 := Student_Nested{
		id: 123,
		name: "Ryan",
		info: StudentInfo{
			class: "Vt6",
			email: "ryan@gmail.com",
			age: 27,
		},
	}

	fmt.Println(student01)

	student02 := Student_Nested{
		456,
		"Robin",
		StudentInfo{
			"Vt6",
			"ryan@gmail.com",
			27,
		},
	}
	fmt.Println(student02)

	// compare 2 struct: 2 struct chi co the so sanh duoc khi cac fields ben trong co the so sanh duoc
	type struct1 struct {
		id int
		name string
	}
	s1 := struct1{1, "A"}
	s2 := struct1{1, "A"}
	if s1 == s2 {
		fmt.Println("s1 == s2")
	} else {
		fmt.Println("s1 != s2")
	}
      // vi du truong hop khong the so sanh
	type struct2 struct {
		id int
		name string
		info map[int] int
	}
/*	s3 := struct2{1, "A", map[int]int{0: 1}}
	s4 := struct2{1, "A", map[int]int{0: 1}}
	  // ko the so sanh s3 voi s4 do trong struct co kieu du lieu map (map ko the so sanh)
	if s3 == s4 {    // invalid operation: s3 == s4 (struct containing map[int]int cannot be compared)
		fmt.Println("s1 == s2")
	} else {
		fmt.Println("s1 != s2")
	}*/

    // zero value cua struct --> la zero value cua cac fields
	var student03 = Student{}
	fmt.Println(student03)

	fmt.Println("=================== Method ================")
	student04 := Student{123, "Robin"}
	name := student04.getName()
	fmt.Println(name)

	fmt.Printf("p1 = %p", &student04)
	fmt.Println()
	student04.changeName()
	fmt.Println(student04.name)

	student04.changeName2()
	fmt.Println(student04.name)

}

// Define method
// func (t Type) methodName(params) returns { // body code}
// (t Type) --> Receiver
// 1. Value Receiver
// 2. Pointer Receiver

// 1. Value Receiver --> ko lam thay doi gia tri cua fields trong struct
func (s Student) changeName()  {
	fmt.Printf("p2 = %p", &s)
	fmt.Println()
	s.name = "Bill"
}

// 2. Pointer Receiver
func (s *Student) changeName2()  {
	s.name = "Bill"

}
func (s Student) getName() string {
	return s.name
}