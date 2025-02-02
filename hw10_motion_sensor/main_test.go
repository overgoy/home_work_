package main

import (
	"testing"
)

func TestDataProcessor(t *testing.T) {
	dataChan := make(chan float64, 10)
	resultChan := make(chan float64, 10)

	go func() {
		dataChan <- 10
		dataChan <- 20
		dataChan <- 30
		dataChan <- 40
		dataChan <- 50
		dataChan <- 60
		dataChan <- 70
		dataChan <- 80
		dataChan <- 90
		dataChan <- 100
		close(dataChan)
	}()

	go dataProcessor(dataChan, resultChan)

	avg := <-resultChan
	if avg != 55 {
		t.Errorf("Expected 55, got %f", avg)
	}
}
