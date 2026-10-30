package main

import "fmt"

func main() {
	const limit int = 4

	for i := 1; i <= limit; i++ {
		for j := 1; j <= limit; j++ {
			fmt.Printf("%dx%d= %2d |", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}
