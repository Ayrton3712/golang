// Multiple assignment
package main

import "fmt"

func main() {
	query, limit, offset := "bat", 10, 0

	query, limit, offset = "ball", offset, limit

	fmt.Printf("%v %v %v\n", query, limit, offset)
}
