// Задание 2: Fan-In - слить два канала
//
// Создай две функции-генератора:
//   - evenNumbers() <-chan int  - генерирует чётные числа 2, 4, 6, 8, 10
//   - oddNumbers() <-chan int   - генерирует нечётные числа 1, 3, 5, 7, 9
//
// Каждая функция запускает горутину, которая кладёт числа в канал и закрывает его.
//
// Напиши функцию merge(ch1, ch2 <-chan int) <-chan int,
// которая читает из обоих каналов и возвращает один объединённый канал.
//
// В main() слей оба генератора и выведи все числа в одну строку.
//
// Ожидаемый вывод (порядок может быть любым):
//   2 1 4 3 6 5 8 7 10 9
//   (или любой другой порядок - главное все 10 чисел)
//
// Запусти: go run main.go

package main

import (
	"fmt"
	"sync"
)

// TODO: напиши функцию evenNumbers() <-chan int

func evenNumbers() <-chan int {
	resuluts := make(chan int)
	go func() {
		defer close(resuluts)
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				resuluts <- i
			}
		}
	}()
	return resuluts
}

// TODO: напиши функцию oddNumbers() <-chan int

func oddNumbers() <-chan int {
	resuluts := make(chan int)
	go func() {
		defer close(resuluts)
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				resuluts <- i
			}
		}
	}()
	return resuluts
}

// TODO: напиши функцию merge(ch1, ch2 <-chan int) <-chan int
// Подсказка: используй WaitGroup и отдельную горутину для закрытия merged

func merge(ch1, ch2 <-chan int) <-chan int {
	merged := make(chan int) // INFO: 	А если буфферезирванный канал?
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range ch1 {
			merged <- v
		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch2 {
			merged <- v
		}
	}()

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged

}

func main() {
	// TODO: создай два канала через генераторы
	ch1 := evenNumbers()
	ch2 := oddNumbers()
	// TODO: слей их через merge
	merged := merge(ch1, ch2)

	// TODO: выведи все числа в одну строку
	for v := range merged {
		fmt.Println(v)
	}
}
