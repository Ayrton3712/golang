package main

import "fmt"

func incrementValue(n int) {
	n++
}

func incrementPointer(n *int) {
	(*n)++ // *n++ is also correct
}

func main() {
	counter := 10

	incrementValue(counter)
	fmt.Printf("incrementValue counter: %d\n", counter)

	incrementPointer(&counter)
	fmt.Printf("incrementPointer counter: %d\n", counter)
}
