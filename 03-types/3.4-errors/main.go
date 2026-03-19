package main

import (
	"errors"
	"fmt"
)

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("На 0 делить нельзя")
	}
	return a / b, nil
}

func main() {
	res, err := safeDivide(1, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err) // На ноль делить нельзя
	} else {
		fmt.Printf("%d\n", res)
	}

	res2, err2 := safeDivide(1, 1)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Printf("%d\n", res2) // 1
	}
}
