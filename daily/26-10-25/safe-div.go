package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func main() {
	quotient1, err1 := divide(10, 2)
	if err1 != nil {
		fmt.Println("Error 1:", err1)
	} else {
		fmt.Println("Quotient 1:", quotient1)
	}

	quotient2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error 2:", err2)
	} else {
		fmt.Println("Quotient 2:", quotient2)
	}
}
