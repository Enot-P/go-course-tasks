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

func CreateUser(name string) error {
	if name == "" {
		return errors.New("Имя не должно быть пустым")
	}
	fmt.Printf("\n%s - создан\n", name)
	return nil
}

var ErrOutOfStock = errors.New("out of stock")

func buyItem(count int) error {
	if count == 0 {
		return ErrOutOfStock
	}

	return nil
}

func readFileMock() error {
	return errors.New("base error")
}

func loadData() error {
	err := readFileMock()
	return fmt.Errorf("%w", err)
}

func main() {
	// task 3.4.1
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

	// task 3.4.2
	errorUser := CreateUser("")         //
	fmt.Println(errorUser)              // Имя не должно быть пустым
	notErrorUser := CreateUser("Anton") // Anton - создан
	fmt.Println(notErrorUser)           // <nil>

	// task 3.4.3
	errCount := buyItem(0)
	if errCount != nil {
		if errors.Is(errCount, ErrOutOfStock) {
			fmt.Printf("\nErrOutOfStock Error: %v\n", errCount) // ErrOutOfStock Error: out of stock
		}
		fmt.Printf("usual error: %v\n", errCount) // usual error: out of stock

	}

	// task 3.4.4
	fmt.Println()
	fmt.Println(loadData()) // base error
}
