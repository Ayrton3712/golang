package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x, y := rand.Intn(10), rand.Intn(10)

	fmt.Println("Before swap:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	x, y = y, x

	fmt.Println("After swap:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
