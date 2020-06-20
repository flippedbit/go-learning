package main

import "fmt"

type cat struct {
	x int
	y int
}

func (c cat) speak() {
	fmt.Println("meow")
}

type dog struct {
	x int
	y int
}

func (d dog) speak() {
	fmt.Println("woof")
}

// only structs with speak() method are included so both cat and dog
type Animal interface {
	speak()
}

// general function allowing Animal interface parameter
// can be used as a general function to do things for all
// structs included
func general(a Animal) {
	fmt.Println("general message")
}

// function taking the Animal interface
// using a switch to differeniate between types
func food(a Animal) {
	switch a.(type) {
	case dog:
		fmt.Println("dog food")
	case cat:
		fmt.Println("cat food")
	}
}

// empty interface type includes all structs
// so any struct will now be able to be passed as a Printer type
type Printer interface{}

func print(p Printer) {
	fmt.Println(p)
}

// only includes doc struct because it is the only struct
// with both speak() and move() methods
// but move() must take two int parameters
type Random interface {
	speak()
	move(x, y int)
}

// does not classify has Random interface
//func (d *dog) move() {
//	d.x++
//	d.y++
//}

func (d *dog) move(x, y int) {
	d.x += x
	d.y += y
}

func main() {
	// dog object
	d1 := dog{x: 1, y: 0}
	// cat object
	c1 := cat{x: 7, y: 10}

	// create slice of Animal interface objects with cat and dog objects
	animals := []Animal{&d1, &c1}
	// loop through Animals slice and call the objects speak() method
	for _, a := range animals {
		a.speak()
	}

	// general function that calls Animal interface
	general(d1)

	// function calling Animal interface using a switch on types
	food(d1)
	food(c1)

	// function using an empty interface to cover all structs
	print(d1)
	print(c1)

	// other interface with more specific methods required
	r := []Random{&d1}
	r[0].move(5, 0)
	print(r[0])
}
