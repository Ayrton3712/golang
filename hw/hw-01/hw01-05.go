/*
  (20 points)
  Write a Go program which declares three variables according
  to the following table and prints the values of the variables
  to the screen.

   Name  |  Type    |  Initial Value
 -------------------------------------
   price      |  int     |  30
   brand      |  string  |
   discount   |          |  true
*/

package main

import "fmt"

func main() {
	price, brand, discount := 30, "some brand", true

	fmt.Println("price:", price, "| brand:", brand, "| discount:", discount)
}
