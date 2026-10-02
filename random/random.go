package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// generate a number n, n = 0, 2, 4, 6, 8
	fmt.Println(rand.Intn(5) * 2)
}
