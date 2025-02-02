package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"sync"
	"time"
)

func sensorDataGenerator(dataChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	timeout := time.After(60 * time.Second)

	for {
		select {
		case <-timeout:
			close(dataChan)
			return
		case <-ticker.C:
			data := generateSecureRandomFloat64()
			select {
			case dataChan <- data:
			default:
			}
		}
	}
}

func generateSecureRandomFloat64() float64 {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0
	}
	return float64(binary.LittleEndian.Uint64(b[:])) / (1 << 64) * 100
}

func dataProcessor(dataChan <-chan float64, resultChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum float64
	count := 0

	for value := range dataChan {
		sum += value
		count++

		if count == 10 {
			resultChan <- sum / 10
			sum = 0
			count = 0
		}
	}

	close(resultChan)
}

func main() {
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

	for avg := range resultChan {
		fmt.Printf("Среднее значение: %.2f\n", avg)
	}
}
