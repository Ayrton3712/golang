package main

import "fmt"

func rectangle(width, height float64) (float64, float64) {
	return width * height, 2 * (width + height)
}

func main() {
	width := 5.5
	height := 3.0

	area, perimeter := rectangle(width, height)

	fmt.Printf("Area: %.2f\nPerimeter: %.2f", area, perimeter)
}
