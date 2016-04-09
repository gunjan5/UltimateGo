// Create a program that declares two anonymous functions. One that counts up to
// 100 from 0 and one that counts down to 0 from 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

import "runtime"
import "fmt"
import "sync"

import (
	"runtime"
)

// Add imports.

// init is called prior to main.
func init() {
	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(2)
}

// main is the entry point for the application.
func main() {
	// Declare a wait group and set the count to two.
	var wg sync.WaitGroup
	wg.Add(2)

	// Declare an anonymous function and create a goroutine.
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("A: ", i)
		}
		// Declare a loop that counts down from 100 to 0 and
		// display each value.

		wg.Done()

		// Decrements the count of the wait group.
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		for i := 100; i > 0; i-- {
			fmt.Println("B: ", i)
		}
		// Declare a loop that counts down from 100 to 0 and
		// display each value.

		wg.Done()

		// Decrements the count of the wait group.
	}()

	wg.Wait()

	// Wait for the goroutines to finish.
}
