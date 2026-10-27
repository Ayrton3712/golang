package main

import "fmt"

func main() {
	// Part 1
	switch temperature := 28; {
	case temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature <= 19:
		fmt.Println("Cool")
	case temperature >= 20 && temperature <= 29:
		fmt.Println("Warm")
	case temperature >= 30:
		fmt.Println("Hot")
	}

	// Part 2
	switch category := 3; category {
	case 1, 2:
		fmt.Println("Wear a jacket")
	case 3, 4:
		fmt.Println("Wear light clothing")
	default:
		fmt.Println("Invalid category")
	}
}
