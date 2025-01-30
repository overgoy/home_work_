package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{
			[]int{1, 3, 5, 7, 9, 11, 13, 15}, // Массив отсортирован
			7,                                // Элемент для поиска
			3,                                // Ожидаемый индекс (7 находится в индексе 3)
		},
		{
			[]int{1, 3, 5, 7, 9, 11, 13, 15},
			9,
			4,
		},
		{
			[]int{1, 3, 5, 7, 9, 11, 13, 15},
			8,
			-1, // Элемента 8 нет в массиве
		},
		{
			[]int{}, // Пустой массив
			5,
			-1, // Пустой массив, ничего не найдено
		},
	}

	for _, test := range tests {
		result := binarySearch(test.arr, test.target)
		if result != test.expected {
			t.Errorf("Для массива %v и целевого значения %d ожидался индекс %d, но получено %d",
				test.arr, test.target, test.expected, result)
		}
	}
}
