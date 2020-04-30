package main

func main() {
	/*
			var a int
			for a < 5 {
				println("a", a)
				a++
			}

		for b := 0; b < 5; b++ {
			println("b", b)
		}
	*/
	var c int
	for {
		if c == 5 {
			break
		}
		println("c", c)
		c++
		if c == 3 {
			continue
		}
		println("continuing...")
	}

	// can range over slices/maps
	s := []string{"First", "Second", "Third"}
	for i, value := range s {
		println(i, value)
	}
}
