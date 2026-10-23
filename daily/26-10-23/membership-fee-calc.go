package main

import "fmt"

func main() {
	const (
		Basic = iota + 1
		Standard
		Premium
	)

	const BaseFee = 10.0
	tier := Premium

	fee := BaseFee + (float64(tier) * 2.5)

	fmt.Printf("Membership tier: %v\nMembership fee: %.2f", tier, fee)
}
