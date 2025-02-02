package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			"hello world hello",
			map[string]int{"hello": 2, "world": 1},
		},
		{
			"Привет, мир! Привет, Go.",
			map[string]int{"привет,": 2, "мир!": 1, "go.": 1},
		},
		{
			"Go, go, GO! go...",
			map[string]int{"go,": 2, "go!": 1, "go...": 1},
		},
		{
			"",
			map[string]int{},
		},
		{
			"123 123 456",
			map[string]int{"123": 2, "456": 1},
		},
	}

	for _, test := range tests {
		result := countWords(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Для ввода '%s' ожидалось %v, но получено %v", test.input, test.expected, result)
		}
	}
}
