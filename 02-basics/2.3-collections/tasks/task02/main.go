// Задание 2: Инвертировать словарь
//
// Напиши функцию invertMap, которая:
//   - принимает map[string]int (название -> число)
//   - возвращает map[int]string (число -> название)
//
// Потом выведи результат: отсортируй ключи по возрастанию и выведи каждую пару.
//
// Используй пакет slices для сортировки ключей.
//
// Ожидаемый вывод:
//   1 -> яблоко
//   2 -> банан
//   3 -> апельсин
//
// Запусти: go run main.go

package main

import (
	"fmt"
	"maps"
	"slices"
)

// TODO: напиши функцию invertMap(m map[string]int) map[int]string

// NOTE: maps.Collect как работает и что за iter.Seq2 чудище такое?
func invertMap(m map[string]int) map[int]string {
	newMap := map[int]string{}

	for key, value := range m {
		newMap[value] = key
	}

	return newMap
}

func main() {
	fruits := map[string]int{
		"яблоко":   1,
		"банан":    2,
		"апельсин": 3,
	}

	// TODO: вызови invertMap и сохрани результат
	inverted := invertMap(fruits)

	// TODO: собери ключи из inverted в срез, отсортируй их
	// и выведи каждую пару в формате: "1 -> яблоко"

	keys := slices.Collect(maps.Keys(inverted))
	slices.Sort(keys)

	for _, v := range keys {
		fmt.Printf("%d -> %v \n", v, inverted[v])
	}
}
