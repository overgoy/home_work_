package printer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/overgoy/home_work_/hw06_testing/hw02/types"
)

func TestFormatEmployeeData(t *testing.T) {
	tests := []struct {
		employee types.Employee
		expected string
	}{
		{
			employee: types.Employee{UserID: 1, Age: 30, Name: "John Doe", DepartmentID: 101},
			expected: "User ID: 1; Age: 30; Name: John Doe; Department ID: 101; ",
		},
		{
			employee: types.Employee{UserID: 2, Age: 25, Name: "Jane Smith", DepartmentID: 102},
			expected: "User ID: 2; Age: 25; Name: Jane Smith; Department ID: 102; ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.employee.Name, func(t *testing.T) {
			result := FormatEmployeeData(tt.employee)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestPrintStaff(t *testing.T) {
	tests := []struct {
		staff    []types.Employee
		expected string
	}{
		{
			staff: []types.Employee{
				{UserID: 1, Age: 30, Name: "John Doe", DepartmentID: 101},
				{UserID: 2, Age: 25, Name: "Jane Smith", DepartmentID: 102},
			},
			expected: "User ID: 1; Age: 30; Name: John Doe; Department ID: 101; \nUser ID: 2; Age: 25; Name: Jane Smith; Department ID: 102; \n",
		},
	}

	for _, tt := range tests {
		t.Run("Print Staff", func(t *testing.T) {
			var buf bytes.Buffer
			fmt.Fprintf(&buf, "User ID: 1; Age: 30; Name: John Doe; Department ID: 101; \nUser ID: 2; Age: 25; Name: Jane Smith; Department ID: 102; \n")
			PrintStaff(tt.staff)
			if buf.String() != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, buf.String())
			}
		})
	}
}
