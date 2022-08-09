package main

import (
	"fmt"
	"sync"
	"time"
)

/*
* An example about the race conditions
*
 */

// A race condition occurs when two or more operations
// must execute in the correct order, but the program
// has not been written so thtat this order is guaranteed
// to be maintained

// func main() {
// 	var data int
// 	go func() {
// 		data++
// 	}()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("hello")
// 	if data == 0 {
// 		fmt.Printf("the value is %v \n", data)
// 	}
// 	fmt.Println(data)
// }

// Memory access Synchronization
// Let's say we have a data race: two concurrent processes
// are attempting to access the same area of the memory,
// and the way they are accessing the memory is not atomic
// Our previous example of a simple data race will do nicely
// with a few modifications

// func main() {
// 	var memoryAccess sync.Mutex
// 	var data int

// 	go func() {
// 		memoryAccess.Lock()
// 		data++
// 		memoryAccess.Unlock()
// 	}()

// 	memoryAccess.Lock()
// 	if data == 0 {
// 		fmt.Println("The value is 0")
// 	} else {
// 		fmt.Printf("The value is %v \n", data)
// 	}
// 	memoryAccess.Unlock()
// }

// IMPLEMENTATION AND UNDERSTANDING OF DEADLOCKS

type value struct {
	mu    sync.Mutex
	value int
}

func main() {

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()

	// page 35
}
