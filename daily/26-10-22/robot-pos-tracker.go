package main

import "fmt"

func main() {
	position := 0
	velocity := 2

	position, velocity = position+velocity, velocity+1

	fmt.Printf("Second 1: Position: %d, Velocity: %d\n", position, velocity)

	position, velocity = position+velocity, velocity+1

	fmt.Printf("Second 2: Position: %d, Velocity: %d\n", position, velocity)

	position, velocity = position+velocity, velocity+1

	fmt.Printf("Second 3: Position: %d, Velocity: %d\n", position, velocity)

	position, velocity = position+velocity, velocity+1

	fmt.Printf("Second 4: Position: %d, Velocity: %d\n", position, velocity)
}
