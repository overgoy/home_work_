package main

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

func sensorDataGenerator(ctx context.Context, dataChan chan<- float64) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			close(dataChan)
			return
		case <-ticker.C:
			dataChan <- generateSecureRandomFloat64()
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
			resultChan <- calculateAverage(buffer)
			buffer = buffer[:0]
		}
	}

	if len(buffer) > 0 {
		resultChan <- calculateAverage(buffer)
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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	dataChan := make(chan float64)
	resultChan := make(chan float64)

	go sensorDataGenerator(ctx, dataChan)
	go dataProcessor(dataChan, resultChan)

	for avg := range resultChan {
		fmt.Printf("Среднее значение: %.2f\n", avg)
	}
}
