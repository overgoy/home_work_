package main

import (
	"math"
	"testing"
)

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name      string
		shape     any
		expected  float64
		expectErr bool
	}{
		{
			name:      "Circle with radius 5",
			shape:     Circle{Radius: 5},
			expected:  math.Pi * 5 * 5,
			expectErr: false,
		},
		{
			name:      "Rectangle with width 10 and height 5",
			shape:     Rectangle{Width: 10, Height: 5},
			expected:  10 * 5,
			expectErr: false,
		},
		{
			name:      "Triangle with base 8 and height 6",
			shape:     Triangle{Base: 8, Height: 6},
			expected:  0.5 * 8 * 6,
			expectErr: false,
		},
		{
			name:      "Invalid shape (string)",
			shape:     "not a shape",
			expected:  0,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateArea(tt.shape)
			if (err != nil) != tt.expectErr {
				t.Errorf("calculateArea() error = %v, wantErr %v", err, tt.expectErr)
			}
			if !tt.expectErr && got != tt.expected {
				t.Errorf("calculateArea() = %v, want %v", got, tt.expected)
			}
		})
	}
}
