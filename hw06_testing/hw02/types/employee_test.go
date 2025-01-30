// types/types_test.go
package types

import (
	"testing"
)

func TestEmployee_String(t *testing.T) {
	tests := []struct {
		employee Employee
		expected string
	}{
		{
			employee: Employee{
				UserID: 1, Age: 30, Name: "John Doe", DepartmentID: 101,
			},
			expected: "User ID: 1; Age: 30; Name: John Doe; Department ID: 101; ",
		},
		{
			employee: Employee{
				UserID: 2, Age: 25, Name: "Jane Smith", DepartmentID: 102,
			},
			expected: "User ID: 2; Age: 25; Name: Jane Smith; Department ID: 102; ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.employee.Name, func(t *testing.T) {
			result := tt.employee.String()
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
