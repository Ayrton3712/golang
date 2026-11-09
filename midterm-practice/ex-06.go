package main

import "fmt"

func add5Value(count int) {
	count += 5
	fmt.Println("Inside add5Value:", count)
}

func add5Point(count *int) {
	*count += 5
	fmt.Printf("%#v\n%d\n", count, *count)
}

func main() {
	var count int
	add5Value(count)
	fmt.Println("After add5Value:", count)
	add5Point(&count)
	fmt.Println("After add5Point:", count)
}
