package main

import "fmt"

func main() {
	price := 12.5
	quantity := 3
	discount := 5.0

	// Calculate subtotal
	subtotal := price * float64(quantity)

	fmt.Printf("Subtotal: %.2f\n", subtotal)

	// Apply discount
	subtotal -= discount

	fmt.Printf("Discounted subtotal: %.2f\n", subtotal)

	// Calculate tax
	tax := subtotal * 0.1

	fmt.Printf("Tax: %.2f\n", tax)

	// Apply tax
	subtotal += tax

	fmt.Printf("Final total: %.2f\n", subtotal)
}
