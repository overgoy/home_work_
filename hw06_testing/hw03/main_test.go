package hw03

import (
	"strconv"
	"testing"
)

func TestGenerateBoard(t *testing.T) {
	tests := []struct {
		size     int
		expected string
	}{
		{
			size:     1,
			expected: "#\n",
		},
		{
			size:     2,
			expected: "# \n #\n",
		},
		{
			size:     3,
			expected: "# #\n # \n# #\n",
		},
		{
			size:     4,
			expected: "# # \n # #\n# # \n # #\n",
		},
		{
			size:     5,
			expected: "# # #\n # # \n# # #\n # # \n# # #\n",
		},
	}

	for _, tt := range tests {
		t.Run("Test_size_"+strconv.Itoa(tt.size), func(t *testing.T) {
			actual := GenerateBoard(tt.size)

			if actual != tt.expected {
				t.Errorf("For size %d, expected %q, but got %q", tt.size, tt.expected, actual)
			}
		})
	}
}
