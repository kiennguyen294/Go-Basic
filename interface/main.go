package main

import "fmt"

// Interface
 /*Interface la mot tap cac method ma cac doi tuong object co the implement
 Interface dinh nghia cac hanh vi cua doi tuong
 Interface the hien tinh da hinh cua doi tuong, moi doi tuong se
 implement method cua interface theo dinh hinh cua doi tuong do*/

 // Khai bao
type Animal interface {
	speak()
}

// multiple interface
type Movement interface {
	move()
}

// Embed interface
type NextAnimal interface {
	Movement
	Animal
}

// Emty interface
func goout(i interface{}) {
	fmt.Println(i)
}
type Dog struct {}

func (d Dog) speak() {
	fmt.Println("woaww woaww")
}

func (d Dog) move() {
	fmt.Println("dog chay bang 4 chan")

}
func main() {
	var animal Animal
	animal = Dog{}

	animal.speak()

	dog := Dog{}
	var m Movement = dog
	m.move()

	var a Animal = dog
	a.speak()

	// su dung nextanimal
	var na NextAnimal = dog
	na.speak()
	na.move()

	// su dung emty interface
	goout(10)
	goout(10.112)
}
