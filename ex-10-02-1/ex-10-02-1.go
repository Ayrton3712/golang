package main

import "fmt"

func main() {
	name1 := "John"
	name2 := "Kevin"
	score1 := 90

	fmt.Println("Hello, " + name1 + " and " + name2)
	fmt.Printf("Hello, %s and %s\n", name1, name2)
	// fmt.Println(name1 + "'s score: " + )
	fmt.Printf("%s's score: %d\n", name1, score1)
}
