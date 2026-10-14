package main

import "fmt"

func main() {
	const pi float64 = 3.14159
	fmt.Printf("pi: %v\n", pi)

	const (
		Jan = 10
		Feb
		Mar = iota
		Apr
		May
	)

	fmt.Printf("Jan: %d\n", Jan)
	fmt.Printf("Feb: %d\n", Feb)
	fmt.Printf("Mar: %d\n", Mar)
	fmt.Printf("Apr: %d\n", Apr)
	fmt.Printf("May: %d\n", May)
}
