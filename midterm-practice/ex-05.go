// Pointers and dereferencing
package main

import (
	"fmt"
	"time"
)

func main() {
	// Part 1: Pointer declarations
	var count1 *int
	count2 := new(int)

	countTemp := 5
	var count3 *int = &countTemp

	t := &time.Time{}

	fmt.Printf("count1: %#v\ncount2: %#v\ncount3: %#v\nt: %#v\n", count1, count2, count3, t)

	// Part 2: Dereferencing
	if count1 != nil {
		fmt.Printf("count1: %v\n", *count1)
	} else {
		fmt.Println("count1 cannot be dereferenced.")
	}

	if count2 != nil {
		fmt.Printf("count2: %v\n", *count2)
	} else {
		fmt.Println("count2 cannot be dereferenced.")
	}

	if count3 != nil {
		fmt.Printf("count3: %v\n", *count3)
	} else {
		fmt.Println("count3 cannot be dereferenced.")
	}

	if t != nil {
		fmt.Printf("t: %v\n", *t)
	} else {
		fmt.Println("t cannot be dereferenced.")
	}

	fmt.Println(t.String())
}
