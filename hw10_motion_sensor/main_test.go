package main

import (
	"sync"
	"testing"
	"time"
)

func TestGenerateSecureRandomFloat64(t *testing.T) {
	value := generateSecureRandomFloat64()
	if value < 0 || value > 100 {
		t.Errorf("Ожидалось число в диапазоне [0, 100], но получено: %f", value)
	}
}

func TestDataProcessor(t *testing.T) {
	dataChan := make(chan float64, 10)
	resultChan := make(chan float64, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go dataProcessor(dataChan, resultChan, &wg)

	for i := 1.0; i <= 10; i++ {
		dataChan <- i
	}
	close(dataChan)

	wg.Wait()
	close(resultChan)

	avg, ok := <-resultChan
	if !ok {
		t.Fatal("Канал обработанных данных был закрыт раньше времени")
	}

	expectedAvg := 5.5
	if avg != expectedAvg {
		t.Errorf("Ожидалось %.2f, но получено %.2f", expectedAvg, avg)
	}
}

func TestFullPipeline(t *testing.T) {
	dataChan := make(chan float64, 10)
	resultChan := make(chan float64, 10)
	var wg sync.WaitGroup

	wg.Add(2)
	go sensorDataGenerator(dataChan, &wg)
	go dataProcessor(dataChan, resultChan, &wg)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	receivedResults := []float64{}
	timeout := time.After(65 * time.Second) // Даем время на выполнение всех операций

loop:
	for {
		select {
		case avg, ok := <-resultChan:
			if !ok {
				break loop
			}
			receivedResults = append(receivedResults, avg)
		case <-timeout:
			t.Fatal("Тест превысил лимит времени")
		}
	}

	if len(receivedResults) == 0 {
		t.Fatal("Ожидалось хотя бы одно среднее значение, но ничего не получено")
	}
}
