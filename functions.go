package main

import "fmt"

var ID int = 0

func main() {
	var id int = setID()
	fmt.Println(id)

	var lowest int = getLowest(10, 4, 103, 11, 19, 125121)
	fmt.Println(lowest)

	updateID(&id)
	fmt.Println(id)
}

// function takes no parameters but increases ID variable and returns it
func setID() int {
	ID++
	return ID
}

// variadic function, can accept unknown amount of parameters of same type so variable amounts are able to change
func getLowest(items ...int) int {
	lowest := items[0]
	for _, i := range items {
		if i < lowest {
			lowest = i
		}
	}
	return lowest
}

// function accepts pointer reference and updates variable memory space instead
func updateID(i *int) {
	*i++
}
