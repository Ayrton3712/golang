// Restaurant bill - Operators
package main

import "fmt"

func main() {
	var total float64

	var subtotal float64 = 2 * 13

	fmt.Printf("Subtotal: %.2f\n", subtotal)

	subtotal += 4 * 2.25

	fmt.Printf("Subtotal: %.2f\n", subtotal)

	subtotal -= 5

	fmt.Printf("Subtotal: %.2f\n", subtotal)

	tip := subtotal * 0.1

	fmt.Printf("Tip: %.2f\n", tip)

	total += tip

	total += subtotal

	fmt.Printf("Final total: %.2f\n", total)

	total /= 2

	fmt.Printf("Amount each person pays: %.2f\n", total)

	visitCount := 24

	visitCount++

	if visitCount%5 == 0 {
		fmt.Printf("You get a reward!\n")
	}
}
