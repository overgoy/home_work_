package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func countWords(text string) map[string]int {
	wordCount := make(map[string]int)

	text = strings.ToLower(text)

	re := regexp.MustCompile(`[^\wа-яА-ЯёЁ]+`)
	text = re.ReplaceAllString(text, " ")

	words := strings.Fields(text)
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	fmt.Println("Введите тект:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Читаем строку с ввода

	text := scanner.Text()
	result := countWords(text)

	fmt.Println("Частота слов:", result)
}
