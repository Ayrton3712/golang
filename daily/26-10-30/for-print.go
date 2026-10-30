package main

import "fmt"

func main() {
	for i := 0; ; i++ {
		if i != 10 {
			fmt.Println(i)
			continue
		}
		break
	}
}
