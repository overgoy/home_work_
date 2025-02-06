package main

import (
	"os"
	"testing"
)

func TestAnalyzeLogs(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testlog*.log")
	if err != nil {
		t.Fatalf("Failed to create temp log file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	logData := "[INFO] System started\n[ERROR] Failed to connect\n[DEBUG] Debugging info\n[INFO] User logged in\n"
	_, err = tempFile.WriteString(logData)
	if err != nil {
		t.Fatalf("Failed to write to temp log file: %v", err)
	}
	tempFile.Close()

	stats, err := analyzeLogs(tempFile.Name(), "INFO")
	if err != nil {
		t.Fatalf("analyzeLogs returned an error: %v", err)
	}

	expected := 2
	if stats["INFO"] != expected {
		t.Errorf("Expected %d INFO logs, got %d", expected, stats["INFO"])
	}
}

func TestWriteStats(t *testing.T) {
	stats := map[string]int{"INFO": 2, "ERROR": 1}
	tempFile, err := os.CreateTemp("", "output*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp output file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	err = writeStats(stats, tempFile.Name())
	if err != nil {
		t.Fatalf("writeStats returned an error: %v", err)
	}

	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read temp output file: %v", err)
	}

	expectedContent := "INFO: 2\nERROR: 1\n"
	if string(content) != expectedContent {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedContent, content)
	}
}
