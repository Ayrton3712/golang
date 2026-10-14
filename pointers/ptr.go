// package main

// import "fmt"

// func main() {
// 	var pi1 *int = new(int)

// 	fmt.Printf("var pi1 *int = new(int)\n")
// 	fmt.Printf("  pi1 address: %v\n", pi1)
// 	fmt.Printf("  pi1 value  : %v\n", *pi1)

// 	var i int = 3
// 	var pi2 *int = &i
// 	fmt.Printf("var i int = 3\n")
// 	fmt.Printf("var pi2 *int = &i\n")
// 	fmt.Printf("  pi2 address: %v\n", pi2)
// 	fmt.Printf("  pi2 value  : %v\n", *pi2)

// 	pi1 = pi2

// 	fmt.Printf("pi1 = pi2\n")
// 	fmt.Printf("  after pi1 = pi2\n")
// 	fmt.Printf("  pi1 address: %v\n", pi1)
// 	fmt.Printf("  pi1 value  : %v\n", *pi1)

// 	*pi1 = 5
// 	fmt.Printf("*pi1 = 5\n")
// 	fmt.Printf("  pi1 address: %v\n", pi1)
// 	fmt.Printf("  pi1 value  : %v\n", *pi1)
// 	fmt.Printf("  pi2 address: %v\n", pi2)
// 	fmt.Printf("  pi2 value  : %v\n", *pi2)
// }
