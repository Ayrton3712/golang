package main

import "fmt"

func main() {
	message := "Go臺灣!"

	// Print byte length
	fmt.Println("Byte length:", len(message))

	var total int
	var nonASCII int
	for index, char := range message {
		fmt.Printf("Byte offset: %d Rune: %c\n", index, char)
		total++

		if char > 127 {
			nonASCII++
		}
	}

	fmt.Printf("Total number of runes: %d\n", total)
	fmt.Printf("Non-ASCII characters: %d\n", nonASCII)
}
