package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func (s *Student) AddPoints(points int) {
	s.Score += points
}

func (s Student) Passed() bool {
	return s.Score >= 60
}

func main() {
	students := []Student{
		{Name: "Mei", Score: 55},
		{Name: "Jun", Score: 82},
		{Name: "Lin", Score: 48},
	}

	var passingCount int
	for i := range students {
		// Add 10 points to every student
		students[i].AddPoints(10)

		// Print each student's name, new score, and result of Passed()
		fmt.Printf("%s: %d, Passed: %v\n",
			students[i].Name,
			students[i].Score,
			students[i].Passed(),
		)

		// Count passing students
		if students[i].Passed() {
			passingCount++
		}
	}

	fmt.Printf("Passing students: %d\n", passingCount)
}
