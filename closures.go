package main

import "fmt"

func intID() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

type user struct {
	name string
	id   int
}

func main() {
	nextID := intID()

	u1 := user{name: "michael", id: nextID()}
	u2 := user{name: "larry", id: nextID()}
	u3 := user{name: "chris", id: nextID()}

	fmt.Println(u1, u2, u3)
}
