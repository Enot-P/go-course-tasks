package main

import (
	"fmt"
)

func Echo[T any](v T) T {
	fmt.Printf("Type:%T\n", v)
	return v
}

func main() {
	fmt.Println(Echo(1))
	fmt.Println(Echo(true))
	fmt.Println(Echo("Hi"))
}
