package main

import "fmt"

func main() {
	sum := 0
	// 0 2 4 6 ... 50 51 53 55 .. 99
	for i := 0; i < 100; i += 2 {
		sum += i
		if i == 50 {
			i = 49
		}
	}
	fmt.Printf("sum = %d\n", sum)
}
