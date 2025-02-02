package main

import (
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{10, 20, 30, 40, 50}, 30.0},
		{[]float64{1, 1, 1, 1, 1}, 1.0},
		{[]float64{5, 10, 15}, 10.0},
		{[]float64{}, 0.0},
	}

	for _, test := range tests {
		result := calculateAverage(test.input)
		if result != test.expected {
			t.Errorf("calculateAverage(%v) = %v, ожидается %v", test.input, result, test.expected)
		}
	}
}
