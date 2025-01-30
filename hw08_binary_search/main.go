package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println("Введите элементы массива, разделенные пробелами:")

	var input string
	fmt.Scanln(&input)

	parts := strings.Fields(input)

	arr := make([]int, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Ошибка при конвертации числа:", err)
			return
		}
		arr = append(arr, num)
	}

	sort.Ints(arr)

	fmt.Println("Введите число для поиска:")
	var target int
	fmt.Scanln(&target)

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден в индексе %d\n", target, index)
	} else {
		fmt.Println("Элемент не найден")
	}
}
