// Задание 1: Безопасный счётчик через Mutex
//
// Создай структуру SafeCounter с:
//   - полем mu типа sync.Mutex
//   - полем value типа int
//   - методом Increment() - увеличивает value на 1 (с блокировкой)
//   - методом Value() int - возвращает value (с блокировкой)
//
// Запусти 1000 горутин, каждая вызывает counter.Increment().
// После wg.Wait() выведи финальное значение.
//
// Проверь отсутствие гонок:
//   go run -race main.go
//
// Ожидаемый вывод:
//   Финальный счётчик: 1000
//   (и никаких WARNING: DATA RACE)
//
// Запусти: go run main.go

package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (sf *SafeCounter) Increment() {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	sf.value++
}

func (sf *SafeCounter) Value() int {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	return sf.value
}

func main() {
	// TODO: создай SafeCounter и запусти 1000 горутин
	// Каждая горутина вызывает counter.Increment()
	// После завершения всех горутин выведи counter.Value()

	safeCounter := SafeCounter{
		mu:    sync.Mutex{},
		value: 0,
	}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			safeCounter.Increment()
		}()
	}
	wg.Wait()

	fmt.Println("Финальное значение: ", safeCounter.Value())

}
