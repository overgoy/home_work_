package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	parts := strings.Fields(input)
	var arr []int
	for _, part := range parts {
		num, err := strconv.Atoi(part) // Преобразуем строку в целое число
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
