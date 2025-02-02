package main

import (
	"context"
	"testing"
	"time"
)

func TestSensorDataGenerator(t *testing.T) {
	dataChan := make(chan float64)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go sensorDataGenerator(ctx, dataChan)

	count := 0
	for range dataChan {
		count++
	}

	if count < 1 || count > 3 {
		t.Errorf("Ожидалось от 1 до 3 значений, получено: %d", count)
	}
}

func TestCalculateAverage(t *testing.T) {
	values := []float64{10, 20, 30, 40, 50}
	expected := 30.0
	result := calculateAverage(values)

	if result != expected {
		t.Errorf("Ожидалось %.2f, получено %.2f", expected, result)
	}
}

func TestDataProcessor(t *testing.T) {
	dataChan := make(chan float64, 10)
	resultChan := make(chan float64)

	go func() {
		for i := 1; i <= 10; i++ {
			dataChan <- float64(i)
		}
		close(dataChan)
	}()

	go dataProcessor(dataChan, resultChan)

	result := <-resultChan
	expected := 5.5

	if result != expected {
		t.Errorf("Ожидалось %.2f, получено %.2f", expected, result)
	}
}
