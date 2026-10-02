// Exercise 02
/*
	Declare two integer variables x and y initialized with random values
	between 0 (inclusive) and 10 (exclusive). Print the values of the
	two variables and check if x is greater than y OR x * y equals x + y.
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x, y := rand.Intn(10), rand.Intn(10)

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("x is greater than y or x * y equals x + y:", (x > y) || (x*y == x+y))
	fmt.Println()
}
