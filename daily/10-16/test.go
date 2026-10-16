package main

import "fmt"

func main() {
	x := 10

	if true {
		x := 20 // x is re-declared here but only within this scope, does not affect the other x
		fmt.Println(x)
	}

	fmt.Println(x)
}
