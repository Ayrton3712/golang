/*
  (20 points)
  Write a Go program that declares two integer variables
  x and y initialized with random values between 0 (inclusive)
  and 10 (exclusive). The program prints the values of the
  two variables and check if x is greater than y or x * y
  equals x + y.
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Initializing x and y to random numbers in [0, 10)
	x, y := rand.Intn(10), rand.Intn(10)

	// Printing values of x and y
	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)

	fmt.Println("x is greater than y OR x * y equals x + y:", (x > y) || (x*y == x+y))
}
