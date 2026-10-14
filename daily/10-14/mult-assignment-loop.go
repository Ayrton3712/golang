package main

import "fmt"

func main() {
	a := 3
	b := 7
	c := 11

	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)
	fmt.Printf("c: %d\n", c)

	a, b, c = b, c, a

	numbers := []int{a, b, c}

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Println("Rotated values:")
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)
	fmt.Printf("c: %d\n", c)
	fmt.Printf("Sum: %d\n", sum)
}
