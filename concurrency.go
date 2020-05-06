package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// set up sync wait group in order to wait until all goroutines return
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	// self calling function
	go func() {
		// run sync Done function upon returning
		defer waitGrp.Done()

		// sleep 5s
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	// self calling function
	go func() {
		// run sync Done function upon returning
		defer waitGrp.Done()

		fmt.Println("World")
	}()

	// wait until sync group is done
	waitGrp.Wait()
}
