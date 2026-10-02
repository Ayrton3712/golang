/*
  (20 points)
  Write a Go program that declares two integer variables x and y.
  Both variables are initialized with random values between
  0 (inclusive) and 10 (exclusive). Print the values of x and y
  together with their summation and average (all in integers).
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Initializing x and y with random values
	var x int = rand.Intn(10)
	var y int = rand.Intn(10)

	// Printing summation and average
	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)
	fmt.Println("Summation:", x+y)
	fmt.Println("Average:", (x+y)/2)
}
