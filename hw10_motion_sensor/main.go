package main

import (
	"fmt"
	"math/rand"
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
			data := rand.Float64() * 100
			select {
			case dataChan <- data:
			default:
			}
		}
	}
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
