/*
  (* 20 points *)
  This file contains an incomplete program. Replace <code>
  with comma separated variables such that the program
  outputs:
  12, 7, 7
*/

package main

import "fmt"

func main() {
	var x, y, z int = 2, 5, 7
	// All right side expressions are evaluated first, then the
	// assignments occur from left to right, so y ends up being
	// equal to 7, and z remains 7
	y, x, y = x+z, y+z, x+y
	fmt.Println(x, y, z)
}

