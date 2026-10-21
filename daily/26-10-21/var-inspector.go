package main

import "fmt"

func main() {
	name := "Ayrton"
	var age = 20
	var height float64 = 1.75
	var student bool
	var score int
	var message string

	// Printing each variable's value & type
	fmt.Printf("name = %v, type = %T\n", name, name)
	fmt.Printf("age = %v, type = %T\n", age, age)
	fmt.Printf("height = %v, type = %T\n", height, height)
	fmt.Printf("student = %v, type = %T\n", student, student)
	fmt.Printf("score = %v, type = %T\n", score, score)
	fmt.Printf("message = %v, type = %T\n", message, message)
}
