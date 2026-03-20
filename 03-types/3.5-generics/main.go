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

type Number interface {
	~int | ~float64
}

func Sum[T Number](items []T) T {
	var sum T = 0
	for _, v := range items {
		sum += v
	}
	return sum
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

	// task 3.5.3
	intSlice := []int{1, 2, 3, 4, 5, 6}
	floatSlice := []float64{1.2, 2.7, 3.2, 4.0, 4.5}

	fmt.Println(Sum(intSlice))   // 21
	fmt.Println(Sum(floatSlice)) // 15.600000000000001
}
