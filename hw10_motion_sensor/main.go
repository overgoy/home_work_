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
			dataChan <- data
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

func dataProcessor(dataChan <-chan float64, resultChan chan<- float64) {
	buffer := make([]float64, 0, 10)

	for value := range dataChan {
		buffer = append(buffer, value)

		if len(buffer) == 10 {
			avg := calculateAverage(buffer)
			resultChan <- avg
			buffer = buffer[:0]
		}
	}

	close(resultChan)
}

func calculateAverage(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
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
