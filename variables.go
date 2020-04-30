package main

import "fmt"

func main() {
	// explicit int var declaration
	var a int
	println("a", a)

	// implicit int var delcaration
	b := 0
	println("b", b)

	// same for string
	var str string = "Michael"
	str2 := "Mike"
	println(str, str2)

	// constant cannot be changed
	const pi = 3.14159
	println(pi)

	// array, cannot change size
	var arr [3]int
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	fmt.Println(arr)

	arr2 := [3]int{3, 4, 5}
	fmt.Println(arr2)

	// slice, can change size dynamically
	slice := []int{9, 8, 7}
	slice = append(slice, 6)
	fmt.Println(slice)

	// structs, can associate other values to struct type almost like a class
	type User struct {
		userID    int
		firstName string
		lastName  string
	}
	var user User
	user.userID = 0
	user.firstName = "Ed"
	user.lastName = "Nigma"
	fmt.Println(user)

	u2 := User{
		userID:    1,
		firstName: "Harvey",
		lastName:  "Dent",
	}
	fmt.Println(u2)

}
