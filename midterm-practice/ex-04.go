// Membership classification
package main

import "fmt"

func main() {
	visit := 15

	firstVisit := visit == 1

	returningVisitor := visit != 1

	silverMember := (visit >= 10) && (visit <= 20)

	goldMember := (visit >= 21) && (visit <= 30)

	platinumMember := visit > 30

	fmt.Println("First visit:", firstVisit)
	fmt.Println("Returning visitor:", returningVisitor)
	fmt.Println("Silver member:", silverMember)
	fmt.Println("Gold member:", goldMember)
	fmt.Println("Platinum member:", platinumMember)
}
