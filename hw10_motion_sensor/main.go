package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

func sensorDataGenerator(dataChan chan<- float64) {
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
	_, err := rand.Read(b[:]) // Использование crypto/rand для генерации случайных байт
	if err != nil {
		return 0
	}
	return float64(binary.LittleEndian.Uint64(b[:])) / (1 << 64) * 100
}

func dataProcessor(dataChan <-chan float64, resultChan chan<- float64) {
	var sum float64
	count := 0

	for value := range dataChan {
		sum += value
		count++

		if count == 10 {
			avg := sum / 10
			resultChan <- avg
			sum = 0
			count = 0
		}
	}

	close(resultChan)
}

func main() {
	dataChan := make(chan float64)
	resultChan := make(chan float64)

	go sensorDataGenerator(dataChan)
	go dataProcessor(dataChan, resultChan)

	for avg := range resultChan {
		fmt.Printf("Среднее значение: %.2f\n", avg)
	}
}
