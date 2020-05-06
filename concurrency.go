package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// allows amount of procs available to run in parallel instead of concurrent
	// without this run concurrently, so first goroutine function runs
	// then switches to second goroutine while waiting
	runtime.GOMAXPROCS(2)

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
