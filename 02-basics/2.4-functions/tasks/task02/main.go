// Задание 2: Счётчик с начальным значением
//
// Напиши функцию makeCounter(start int) func() int, которая:
//   - принимает начальное значение
//   - возвращает функцию-счётчик
//   - каждый вызов возвращает следующее число, начиная с start
//
// Ожидаемый вывод:
//   Счётчик от 5:
//   5
//   6
//   7
//   Счётчик от 100:
//   100
//   101
//   Счётчики независимы - счётчик от 5 продолжает:
//   8
//
// Запусти: go run main.go

package main

import "fmt"

// TODO: напиши функцию makeCounter(start int) func() int
// Подсказка: текущее значение храни в переменной внутри makeCounter,
// и обращайся к ней из возвращаемой функции (это и есть замыкание)

func makeCounter(start int) func() int {
	counter := start
	return func() int {
		counter++
		return counter
	}
}

func main() {
	counter := makeCounter(0)
	fmt.Println("Счетчик_1: ", counter())
	fmt.Println("Счетчик_1: ", counter())
	fmt.Println("Счетчик_1: ", counter())
	fmt.Println("Счетчик_1: ", counter())
	fmt.Println("Счетчик_1: ", counter())
	fmt.Println("=============================")
	counter_2 := makeCounter(100)
	fmt.Println("Счетчик_2: ", counter_2())
	fmt.Println("Счетчик_2: ", counter_2())
	fmt.Println("Счетчик_2: ", counter_2())
	fmt.Println("Счетчик_2: ", counter_2())
}
