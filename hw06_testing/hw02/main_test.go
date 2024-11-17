package hw02

import (
	"github.com/overgoy/home_work_/hw06_testing/hw02/printer"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetFilePath(t *testing.T) {
	path := getFilePath()
	assert.NotEmpty(t, path, "Path should not be empty")
}

func TestLoadStaff_Success(t *testing.T) {
	testData := `[
		{"userId": 1, "age": 30, "name": "John Doe", "departmentId": 101},
		{"userId": 2, "age": 25, "name": "Jane Smith", "departmentId": 102}
	]`

	tmpFile, err := os.CreateTemp("", "test_data_*.json")
	if err != nil {
		t.Fatalf("unable to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Удалим файл после теста

	if _, err := tmpFile.Write([]byte(testData)); err != nil {
		t.Fatalf("unable to write test data: %v", err)
	}

	tmpFile.Close()

	staff, err := loadStaff(tmpFile.Name())

	assert.NoError(t, err)
	assert.Len(t, staff, 2)
	assert.Equal(t, "John Doe", staff[0].Name)
	assert.Equal(t, 30, staff[0].Age)
	assert.Equal(t, 101, staff[0].DepartmentID)
}

func TestLoadStaff_Error(t *testing.T) {
	staff, err := loadStaff("non_existent_file.json")

	assert.Error(t, err)
	assert.Nil(t, staff)
}

func TestPrintStaff(t *testing.T) {
	testData := `[
		{"userId": 1, "age": 30, "name": "John Doe", "departmentId": 101}
	]`

	tmpFile, err := os.CreateTemp("", "test_data_*.json")
	if err != nil {
		t.Fatalf("unable to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Удалим файл после теста

	if _, err := tmpFile.Write([]byte(testData)); err != nil {
		t.Fatalf("unable to write test data: %v", err)
	}

	tmpFile.Close()

	staff, err := loadStaff(tmpFile.Name())
	if err != nil {
		t.Fatalf("unable to load staff: %v", err)
	}

	printer.PrintStaff(staff)
}
