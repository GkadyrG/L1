package main

import "fmt"

func main() {
	testCases := []struct {
		arr      []int
		target   int
		expected int
	}{
		// Базовые тесты
		{[]int{1, 3, 5, 7, 9}, 5, 2}, // Элемент в середине
		{[]int{1, 3, 5, 7, 9}, 1, 0}, // Элемент в начале
		{[]int{1, 3, 5, 7, 9}, 9, 4}, // Элемент в конце
		{[]int{1, 3, 5, 7, 9}, 3, 1}, // Элемент слева от центра
		{[]int{1, 3, 5, 7, 9}, 7, 3}, // Элемент справа от центра

		// Элемент не найден
		{[]int{1, 3, 5, 7, 9}, 0, -1},  // Меньше всех
		{[]int{1, 3, 5, 7, 9}, 10, -1}, // Больше всех
		{[]int{1, 3, 5, 7, 9}, 4, -1},  // Между элементами

		// Крайние случаи
		{[]int{}, 5, -1},  // Пустой массив
		{[]int{5}, 5, 0},  // Массив из одного элемента (найден)
		{[]int{5}, 3, -1}, // Массив из одного элемента (не найден)

		// Массив с четным количеством элементов
		{[]int{1, 3, 5, 7}, 3, 1},
		{[]int{1, 3, 5, 7}, 5, 2},

		// Массив с повторяющимися элементами
		{[]int{1, 2, 2, 2, 3}, 2, 2}, // Возвращает любое вхождение
	}

	// Запускаем тесты
	for i, tc := range testCases {
		result := binSearch(tc.arr, tc.target)
		if result == tc.expected {
			fmt.Printf("Тест %d: PASSED\n", i+1)
		} else {
			fmt.Printf("Тест %d: FAILED. Ожидалось %d, получено %d\n",
				i+1, tc.expected, result)
		}
	}
}

func binSearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		m := (l + r) / 2
		if arr[m] == target {
			return m
		} else if arr[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
