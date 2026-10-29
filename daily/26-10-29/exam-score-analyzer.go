package main

import "fmt"

func main() {
	scores := []int{85, 42, 91, 67, 58}

	var passing, sum int
	for i, score := range scores {
		// Print each score with its index, starting at 0
		fmt.Printf("Score %d: %d\n", i, score)

		// Count how many scores are passing
		if score >= 60 {
			passing += 1
		}

		// Calculate sum of all scores
		sum += score
	}

	// Calculate average
	average := float64(sum) / float64((len(scores)))

	// Print passing count and the average score to two decimal places
	fmt.Printf("Passing students: %d\nAverage score: %.2f\n", passing, average)
}
