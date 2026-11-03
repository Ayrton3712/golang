package main

import "fmt"

func main() {
	type Student struct {
		Name  string
		Score int
	}

	students := []Student{
		{Name: "Mei", Score: 82},
		{Name: "Jun", Score: 55},
		{Name: "Lin", Score: 94},
		{Name: "Wei", Score: 71},
	}

	// Declare counts, highestScore (index of highest-scoring student) and sum
	var passingCount, topIndex, sum int
	for i, student := range students {
		// Print every student's name and score
		fmt.Printf("%s: %d\n", student.Name, student.Score)

		// Count students with passing score
		if student.Score >= 60 {
			passingCount++
		}

		// Find student with highest score
		if students[topIndex].Score < student.Score {
			topIndex = i
		}

		// Add scores
		sum += student.Score
	}

	// Calculate average
	averageScore := float64(sum) / float64(len(students))

	// Print
	fmt.Printf("Passing students: %d\n", passingCount)
	fmt.Printf("Top student: %s (%d)\n", students[topIndex].Name, students[topIndex].Score)
	fmt.Printf("Average: %.2f\n", averageScore)
}
