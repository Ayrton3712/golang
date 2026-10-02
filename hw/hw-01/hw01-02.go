/*
  (20 points)
  Write a Go program to declare and initialize in one line
  three variables x, y, and z with initial values 10, 3.14,
  "Hello World" respectively. Print their values.
*/

package main

import "fmt"

func main() {
	x, y, z := 10, 3.14, "Hello World"

	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)
	fmt.Println("Value of z:", z)
}
