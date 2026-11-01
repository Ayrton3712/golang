package main

import "fmt"

func main() {
	counts := map[string]int{
		"Go":     4,
		"Python": 7,
		"C":      5,
		"Java":   3,
	}

	// Add Rust with a count of 6
	counts["Rust"] = 6

	// Increase Go by 2
	counts["Go"] += 2

	// Comma-ok idiom to check if C++ exists
	value, exists := counts["C++"]
	if exists {
		fmt.Println("C++:", value)
	} else {
		fmt.Println("C++ not found")
	}

	var highestLanguage string
	var highestFrequency int
	for language, frequency := range counts {
		if highestFrequency < frequency {
			highestLanguage, highestFrequency = language, frequency
		}
	}

	fmt.Printf("Most frequent: %s (%d)\n", highestLanguage, highestFrequency)
	fmt.Println("Total languages:", len(counts))
}
