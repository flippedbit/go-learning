package main

import "fmt"

type User struct {
	ID int
	FirstName string
	LastName string
}

func main() {
	u1 := User{
		ID: 1,
		FirstName: "Victor",
		LastName: "Fries",
	}
	fmt.Println(u1.ID)
	_, err := u1.updateID(2)
	if err == nil {
		fmt.Println(u1.ID)
	}
}

// if you dont associate type is pointer struct type it will create a new struct type with variables to adjust
func (u *User) updateID(i int) (int, error) {
	if i != u.ID {
		u.ID = i
		return i, nil
	}
	return u.ID, fmt.Errorf("ID cannot be the same")
}