package main

func printStuff() {
	// this line will not be execute until the function returns
	defer println("second line")

	println("third line")
}

func main() {
	println("first line")
	printStuff()
}
