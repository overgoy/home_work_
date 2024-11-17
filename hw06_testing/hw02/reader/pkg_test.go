// hw02/reader/reader_test.go
package reader

import (
	"path/filepath"
	"testing"
)

func TestReadJSON(t *testing.T) {
	tests := []struct {
		fileName    string
		expectErr   bool
		expectCount int
	}{
		{
			fileName:    "testdata/test_data_valid.json", // Путь к файлу с корректными данными
			expectErr:   false,
			expectCount: 2,
		},
		{
			fileName:    "testdata/test_data_invalid.json", // Путь к файлу с некорректными данными
			expectErr:   true,
			expectCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			// Вставляем абсолютный путь к файлу
			absPath, err := filepath.Abs(tt.fileName)
			if err != nil {
				t.Fatalf("could not get absolute path: %v", err)
			}

			staff, err := ReadJSON(absPath)

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}

			if len(staff) != tt.expectCount {
				t.Errorf("expected %d employees, got %d", tt.expectCount, len(staff))
			}
		})
	}
}
