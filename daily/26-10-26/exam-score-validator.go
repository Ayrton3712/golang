package main

import (
	"errors"
	"fmt"
)

func evaluateScore(score int) (string, error) {
	if score < 0 || score > 100 {
		return "", errors.New("invalid score")
	}
	if score >= 90 {
		return "Excellent", nil
	}
	if score >= 60 {
		return "Passed", nil
	}
	return "Failed", nil
}

func main() {
	score := 85

	if result, err := evaluateScore(score); err != nil {
		fmt.Printf("Score %d: Error: %s\n", score, err)
	} else {
		fmt.Printf("Score %d: %s\n", score, result)
	}
}
