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

func IndexOf[T comparable](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

func Values[K comparable, V any](m map[K]V) []V {
	res := make([]V, 0, len(m))

	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func Contains[T comparable](item []T, target T) bool {
	for _, v := range item {
		if v == target {
			return true
		}
	}
	return false
}

type Store[T any] struct {
	items []T
}

func (s *Store[T]) Add(item T) {
	s.items = append(s.items, item)
}

func (s *Store[T]) All() []T {
	return s.items
}

func main() {
	fmt.Println(Echo(1))
	fmt.Println(Echo(true))
	fmt.Println(Echo("Hi"))
	fmt.Println()

	fmt.Println("/////task 3.5.2/////")
	// task 3.5.2
	slice1 := make([]int, 0)
	fmt.Println(First(slice1))
	slice1 = append(slice1, 1)
	fmt.Println(First(slice1))

	fmt.Println("")
	fmt.Println("/////task 3.5.3/////")
	// task 3.5.3
	intSlice := []int{1, 2, 3, 4, 5, 6}
	floatSlice := []float64{1.2, 2.7, 3.2, 4.0, 4.5}

	fmt.Println(Sum(intSlice))   // 21
	fmt.Println(Sum(floatSlice)) // 15.600000000000001
	fmt.Println("")

	fmt.Println("/////task 3.5.4/////")
	// task 3.5.4
	fmt.Println(IndexOf(intSlice, 3))  // 2
	fmt.Println(IndexOf(intSlice, 30)) // -1

	fmt.Println("")
	fmt.Println("/////task 3.5.5/////")

	// task 3.5.5
	firstPair := Pair[string, int]{
		Key:   "Key",
		Value: 101,
	}
	secondPair := Pair[int, string]{
		Key:   11001,
		Value: "Value",
	}
	fmt.Println(firstPair)
	fmt.Println(secondPair)

	fmt.Println("")
	fmt.Println("/////task 3.5.6/////")
	// task 3.5.6
	m := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 7,
	}

	fmt.Println(Values(m))

	fmt.Println("")
	fmt.Println("/////task 3.5.7/////")
	// task 3.5.7
	var storeString Store[string]
	var storeInt Store[int]

	storeString.Add("123")
	storeInt.Add(123456)
	fmt.Println(storeString.All())
	fmt.Println(storeInt.All())
	fmt.Println(Contains(storeString.All(), "123")) // true
	fmt.Println(Contains(storeInt.All(), 123456))   // true
	fmt.Println(Contains(storeString.All(), "12"))  // false
	fmt.Println(Contains(storeInt.All(), 12345))    // false
}
