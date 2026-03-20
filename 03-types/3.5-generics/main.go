package main

import (
	"fmt"
)

func Echo[T any](v T) T {
	fmt.Printf("Type:%T\n", v)
	return v
}

func First[T any](items []T) (T, bool) {
	if len(items) == 0 {
		var zero T
		return zero, false
	}
	return items[0], true
}

func main() {
	fmt.Println(Echo(1))
	fmt.Println(Echo(true))
	fmt.Println(Echo("Hi"))
	fmt.Println()

	// task 3.5.2
	slice1 := make([]int, 0)
	fmt.Println(First(slice1))
	slice1 = append(slice1, 1)
	fmt.Println(First(slice1))
}
