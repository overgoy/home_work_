package main

import (
	"testing"
)

func TestCompareBooksByYear(t *testing.T) {
	tests := []struct {
		book1    *Book
		book2    *Book
		expected bool
	}{
		{
			book1:    NewBook(1, "Go Programming", "John Doe", 2020, 300, 4.5),
			book2:    NewBook(2, "Learning Go", "Jane Smith", 2021, 250, 4.7),
			expected: false, // book2 is newer, so it should be greater
		},
		{
			book1:    NewBook(3, "Advanced Go", "Alice", 2022, 350, 4.8),
			book2:    NewBook(4, "Intro to Go", "Bob", 2020, 200, 4.2),
			expected: true, // book1 is newer, so it should be greater
		},
	}

	for _, tt := range tests {
		t.Run("Year Comparison", func(t *testing.T) {
			actual := CompareBooksByYear(tt.book1, tt.book2)
			if actual != tt.expected {
				t.Errorf("For books %v and %v, expected %v, but got %v", tt.book1, tt.book2, tt.expected, actual)
			}
		})
	}
}

func TestCompareBooksBySize(t *testing.T) {
	tests := []struct {
		book1    *Book
		book2    *Book
		expected bool
	}{
		{
			book1:    NewBook(1, "Go Programming", "John Doe", 2020, 300, 4.5),
			book2:    NewBook(2, "Learning Go", "Jane Smith", 2021, 250, 4.7),
			expected: true,
		},
		{
			book1:    NewBook(3, "Advanced Go", "Alice", 2022, 350, 4.8),
			book2:    NewBook(4, "Intro to Go", "Bob", 2020, 400, 4.2),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run("Size Comparison", func(t *testing.T) {
			actual := CompareBooksBySize(tt.book1, tt.book2)
			if actual != tt.expected {
				t.Errorf("For books %v and %v, expected %v, but got %v", tt.book1, tt.book2, tt.expected, actual)
			}
		})
	}
}

func TestCompareBooksByRate(t *testing.T) {
	tests := []struct {
		book1    *Book
		book2    *Book
		expected bool
	}{
		{
			book1:    NewBook(1, "Go Programming", "John Doe", 2020, 300, 4.5),
			book2:    NewBook(2, "Learning Go", "Jane Smith", 2021, 250, 4.7),
			expected: false,
		},
		{
			book1:    NewBook(3, "Advanced Go", "Alice", 2022, 350, 4.8),
			book2:    NewBook(4, "Intro to Go", "Bob", 2020, 400, 4.2),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run("Rate Comparison", func(t *testing.T) {
			actual := CompareBooksByRate(tt.book1, tt.book2)
			if actual != tt.expected {
				t.Errorf("For books %v and %v, expected %v, but got %v", tt.book1, tt.book2, tt.expected, actual)
			}
		})
	}
}
