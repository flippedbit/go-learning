package main

type HTTPRequest struct {
	Method string
}

func main() {
	a := 5

	// normal if statements
	if a > 1 {
		println("a greater than 1, ", a)
	} else {
		println("a is not greater than 1, ", a)
	}

	// if, else if
	if a > 1 {
		println("a greater than 1, ", a)
	} else if a < 10 {
		println("a less than 10, ", a)
	}

	// simple initialization statement, local variables are declared prior to if block running
	if y, z := 0, 1; y < z {
		println("y is less than z")
	} else {
		println("y is greater than z")
	}

	// switches, implied breaks in each statement
	r := HTTPRequest{Method: "HEAD"}
	switch r.Method {
	case "GET":
		println("Get request")
		// breaks are implied but can be used
		break
	case "POST":
		println("Post request")
		// can also use passthrough to continue after switch catch
	case "DELETE":
		println("Delete request")
	case "PUT":
		println("Put request")
	default:
		// default catch-all case if not defined
		println("Unhandled method")
	}
}
