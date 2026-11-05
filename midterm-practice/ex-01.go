package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	var1, var2, var3 := getConfig()

	fmt.Printf("%v %v %v\n", var1, var2, var3)
}
