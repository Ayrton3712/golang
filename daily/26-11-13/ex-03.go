// Bulls and Cows game
// Guess 4 digits from 0 to 9
// secret: 5 3 1 4, all digits distinct, each digit between 1 and 9
// guess: 1 7 9 4
// answer: 1A1B
// 1A means there's one digit in the correct position
// 1B means there's one digit that appears but in the incorrect position
// another guess: 5 4 1 7
// 2A1B

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Generate secret
	secret := []int{
		rand.Intn(9) + 1,
		rand.Intn(9) + 1,
		rand.Intn(9) + 1,
		rand.Intn(9) + 1,
	}

	// fmt.Println(secret)
	fmt.Println("--- Bulls and Cows ---")

	// Take guess
	guess := []int{0, 0, 0, 0}

	for i := 0; i < len(guess); i++ {
		fmt.Printf("Guess %d: ", i+1)
		fmt.Scan(&guess[i])
	}

	fmt.Println(guess)

	var x, y int
	// Comparing each digit of guess with each digit of secret
	for i := 0; i < len(guess); i++ {
		for j := 0; j < len(secret); j++ {
			if guess[i] == secret[j] && i == j {
				x++
			} else if guess[i] == secret[j] {
				y++
			}
		}
	}

	fmt.Printf("%dA%dB\n", x, y)
}
