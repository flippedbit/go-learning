package main

// creates "parent" Animal struct that includes other composition "child" structs
type Animal struct {
	Bird
	Fish
}

// bird struct and methods to move and fly
type Bird struct{}

func (b Bird) fly() {
	println("flying")
}
func (b Bird) move() {
	println("bird moving")
}

// fish struct and methods to move and swim
type Fish struct{}

func (f Fish) swim() {
	println("swimming")
}
func (f Fish) move() {
	println("fish moving")
}

// a.move() both composition structs have a move function that have the same parameter and return so this is an ambiguous call
// you must create a method for the "parent" struct that overrides the individual struct methods
func (a Animal) move() {
	println("moving")
}

func main() {
	var a Animal
	a.fly()  // the Animal struct passes to the Bird struct and calls the fly() method
	a.swim() // the Animal struct passes to the Fish struct and calls the swim() method
	a.move() // will simply output "moving" from the generic Animal.moving() method
}
