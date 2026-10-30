package main

import "fmt"

func main() {
	scores := []int{70, 80, 90}

	// Append 85 and 95
	scores = append(scores, 85, 95)

	// Assign scores to backup
	backup := scores

	// Change backup[0] to 100
	backup[0] = 100

	// Print both slices
	fmt.Println("Scores:", scores)
	fmt.Println("Backup:", backup)

	// Print the length of scores
	fmt.Println("Length:", len(scores))

	// Calculate and print sum of all scores
	var sum int
	for _, value := range scores {
		sum += value
	}
	fmt.Println("Sum:", sum)
}
