package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func intID() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

func kill_time(i int64) {
	for a := int64(0); a < i; a++ {
		continue
	}
}

func perf_time(f func(int64), a int64) func() time.Duration {
	return func() time.Duration {
		start := time.Now()
		f(a)
		return time.Since(start)
	}
}

type user struct {
	name string
	id   int
}

func main() {
	// if an argument is passed then run the functions to see how long it takes to loop through
	if len(os.Args) > 1 {
		i, _ := strconv.ParseInt(os.Args[1], 10, 64)
		a := perf_time(kill_time, i)
		fmt.Println(a())
	} else {
		// if no argument is passed then show how to run the functions for incrementing userID
		nextID := intID()

		u1 := user{name: "michael", id: nextID()}
		u2 := user{name: "larry", id: nextID()}
		u3 := user{name: "chris", id: nextID()}

		fmt.Println(u1, u2, u3)
	}
}
