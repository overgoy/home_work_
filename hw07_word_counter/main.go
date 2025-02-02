package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(text string) map[string]int {
	wordCount := make(map[string]int)

	text = strings.ToLower(text)

	words := strings.Fields(text)
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	fmt.Println("Введите текст:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	text := scanner.Text()
	result := countWords(text)

	fmt.Println("Частота слов:", result)
}
