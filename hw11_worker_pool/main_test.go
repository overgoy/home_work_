package main

import (
	"sync"
	"testing"
)

func TestCounterIncrement(t *testing.T) {
	counter := &Counter{}
	var wg sync.WaitGroup

	numWorkers := 100
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	if counter.GetValue() != numWorkers {
		t.Errorf("Ожидаемое значение: %d, Полученное: %d", numWorkers, counter.GetValue())
	}
}

func TestConcurrentAccess(t *testing.T) {
	counter := &Counter{}
	var wg sync.WaitGroup

	numWorkers := 1000
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	expected := numWorkers
	actual := counter.GetValue()

	if actual != expected {
		t.Errorf("Ошибка конкурентного доступа. Ожидалось: %d, Получено: %d", expected, actual)
	}
}
