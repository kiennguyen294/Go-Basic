package main

import (
	"fmt"
)

func Chao()  {
	fmt.Println("Hello")
}

func sayHello(name string)  {
	fmt.Printf("Hello", name)
}

func  greeting(name string) string  {
	result := fmt.Sprintf("Hello %s", name)
	return result
}

// Multiple return value
func recInfo(w, h int) (int, int, int) {
	area := w * h
	return w, h, area
}

// Named reture value
func recInfo_Named(w, h int) (weight int, height int, isSquare bool ) {
	isSquare = w == h   // neu w == h thi isSquare = True
	return w, h, isSquare
}
func main()  {

	//var myString string = "Kiên"
	//runes := []rune(myString)
	//
	//for i:=0; i<len(runes); i++ {
	//	fmt.Printf("%c", runes[i])
	//}
	//fmt.Println("\n")
	//var myRune rune = 'A'
	//fmt.Printf("%T", myRune)
	//Chao()
	//sayHello("Ken")
	//result := greeting("Alice")
	//fmt.Println(result)

	w, h, area := recInfo(100, 200)
	fmt.Println("width=", w)
	fmt.Println("height=", h)
	fmt.Println("area=", area)

	w1, h1, isSquare := recInfo_Named(200, 200)
	if isSquare {
		fmt.Println("This is square")
	} else {
		fmt.Println("width=",w1)
		fmt.Println("height=", h1)
	}

	fmt.Println("======== Slice ===============")
	// khai bao slice
	var slice []int
	fmt.Println(slice)

	// khai bao va khoi tao slice
	var slice1 = []int {1,2,3,4}
	fmt.Println(slice1)

	// tao mot slice tu 1 array
	var array = [4]int {1,2,3,4}
	slice2 := array[1:3]  // array[1] --> array[2]
	fmt.Println(slice2)
	slice3 := array[:]
	fmt.Println(slice3)

	slice4 := array[2:]
	fmt.Println(slice4)

	slice5 := array[:3]
	fmt.Println(slice5)

	// tao mot slice tu mot slice
	var slice6 = []int {1,2,3,4,5,6,7,8,9}
	slice7 := slice6
	fmt.Println(slice7)

	slice8 := slice6[1:]
	fmt.Println(slice8)

	// slice => reference type
	var array1 = [5]int {1,2,3,4,5}
	slice9 := array1[:]
	slice9[0] = 777
	fmt.Println(slice9)
	fmt.Println(array1)

	// length va capacity cua slice
	countries := [...]string {"VN", "USA", "CANADA", "FRANCE", "CHINA"}
	slice10 := countries[1:3] // CANADA
	fmt.Println(slice10)
	fmt.Println(len(slice10))
	fmt.Println(cap(slice10))

	// len of slice la so luong phan tu cua slice
	// cap of slice la so luong phan tu cua underlying array bat dau tu vi tri start khi ma slice duoc tao
	// trong vi du tren, slice10 duoc tao tu array bat dau tu phan tu 1 ==> cap la [1,2,3,4] ==> cap = 4
	// underlying array

	// make, copy, append
	///1. make
	fmt.Println("=======================================")
	slice11 := make([]int, 2, 5)
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))

	slice12 := make([]int, 2)
	fmt.Println(slice12)
	fmt.Println(len(slice12))
	fmt.Println(cap(slice12))

	/// 2.append them gia tri vao slice
	fmt.Println("========== append slice =====================")
	var slice13 []int
	slice13 = append(slice13, 100)
	fmt.Println(slice13)

	slice14 := []int{10, 20, 30, 40, 50}
	newslice14 := slice14[1:3]
	newslice14 = append(newslice14, 60)
	fmt.Println(newslice14)
	fmt.Println(slice14)

	newslice15 := append(slice14,70)
	fmt.Println("#############cap of slice 15")
	fmt.Println(cap(newslice15))
	fmt.Println(len(newslice15))

	///3. copy
	src := []string {"A", "B", "C", "D"}
	dest := make([]string, 2)

	number := copy(dest, src)
	fmt.Println(number)

	// delete item with index = 1
	src = append(src[:1], src[2:]...) // noi slice - slice; append(slice1, slice2...)
	fmt.Println(src)

	// variadic function
	fmt.Println("================== Variadic Function =============")
	addItem(1, 100,200, 300, 400)

	var list = []int {1,2,3,4}
	addItem(100, list...)

	change(list...)
	fmt.Println(list)


	// map trong golang
	// key - value // key la duy nhat
	fmt.Println("============== Map ===================")
	// khai bao
	var myMap = make(map[string]int) // kieu du lieu key: string, value: int
	fmt.Println(myMap)
	if myMap == nil {
		fmt.Println("myMap1 = nil")
	} else {
		fmt.Println("myMap != nil")
	}

	var myMap1 map[string]int
	fmt.Println(myMap1)        // zero value cua map la nil
	if myMap1 == nil {
		fmt.Println("myMap1 = nil")
	}

	// khai bao map voi gia tri khoi tao
	myMap2 := map[string]int {
		"key1" : 1,
		"key2" : 2,
		"key3" : 3,
	}
	fmt.Println(myMap2)

	// them 1 phan tu vao map
	myMap2["key4"] = 4
	myMap2["key5"] = 5
	fmt.Println(myMap2)

	// del mot phan tu trong map = delete(map, key)
	delete(myMap2, "key1")
	fmt.Println(myMap2)

	// length cua map
	fmt.Println("len of map ", len(myMap2))

	// map la reference type
	myMap3 := myMap2
	fmt.Println("myMap3 is ", myMap3)
	myMap3["key5"] = 1000
	fmt.Println("myMap2  after change is ", myMap2)

	// truy cap 1 phan tu trong map
	value, found := myMap2["key1000"]
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("ko tim thay gia tri voi key1000")
	}

	// trong map khong co toan tu == --> khong the kiem tra map1 == map2

	fmt.Println("================= Con tro trong Golang ================")
	// pointer la mot bien va no luu tru dia chi vung nho cua bien khac

	//khai bao
	a := 100
	var pointer *int
	pointer = &a
	fmt.Println(pointer)  // print ra dia chi vung nho cua bien a
	fmt.Printf("%T", pointer)

	// cach khai bao con tro t2

	b := 123
	p2 := new(int)  // <==> var p2 *int
	p2 = &b
	fmt.Println()
	fmt.Println(p2)
	fmt.Printf("%T", p2)

	// zero value cua pointer
	var p3 *int   // zero value la nil
	p4 := new(int) // zero value la hexa dia chi o nho
	fmt.Println(p3)
	fmt.Println(p4)

	// can thiep gia tri cua bien
	a1 := 100
	var p5 *int
	p5 = & a1
	*p5 = 999   // sua gia tri cua a1 thanh 999
	fmt.Println(p5)
	fmt.Println(a1)

	// con tro tro toi array

	array10 := [3]int {1,2,3}
	var p6 *[3]int
	p6 = &array10
	fmt.Println(p6)
	fmt.Printf("%T", p6)
	fmt.Println()

	// truyen pointer vao function
	c := 30
	var p7 *int = &c
	applyPointer(p7)
	fmt.Println(c)

	var array21 [3]*string
	array22 := [3]*string {new(string), new(string), new(string)}

	*array22[0] = "Red"
	*array22[1] = "Blue"
	*array22[2] = "Green"

	array21 = array22

	fmt.Println("++++++++++++ Array21 ++++++++++++++" )
	fmt.Println(array21)
	fmt.Println(array22)

	fmt.Println(*array21[0])



}

// variadic function
/// la function ko gioi han dau vao
func addItem(item int, list...int)  {
	// 100, 200, 300, 400 ---new slice----> int[] {100, 200, 300, 400}
	// []int {1,2,3,4} ---new slice----> int[] {[]int {1,2,3,4}} ===> failed
	list = append(list, item)
	fmt.Println(list)
	
}

func change(list ...int)  {
	list[0] = 999

}

// truyen con tro vao ham
func applyPointer(pointer *int)  {
	*pointer = 777
}
