package main

import "fmt"

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	a := 3
	b := 7

	fmt.Printf("before swap:\n")
	fmt.Printf("  a: %d\n", a)
	fmt.Printf("  b: %d\n", b)

	swap(&a, &b)

	fmt.Printf("after swap:\n")
	fmt.Printf("  a: %d\n", a)
	fmt.Printf("  b: %d\n", b)
}
