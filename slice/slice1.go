//Declare a nil slice of integers. Create a loop that appends 10 values
// to the slice. Iterate over the slice and display each value.

package main

import (
	"fmt"
)

func main() {
	var s []int

	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println("Length: ", len(s), " Capacity: ", cap(s))

	}

	// for i, v := range s {
	// 	fmt.Println(i, v)
	// }
}
