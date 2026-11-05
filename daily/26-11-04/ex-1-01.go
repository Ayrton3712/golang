// Write a program to print a random number, between 1 and 5, of stars (*) to the
// console

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(5) + 1 // Generate random number between 1 and 5

	fmt.Println("Random number:", randomNumber)

	// Loop approach
	// for i := 0; i < randomNumber; i++ {
	// 	fmt.Printf("*")
	// }

	// Strings package approach
	stars := strings.Repeat("*", randomNumber)

	fmt.Printf("%s\n", stars)
}
