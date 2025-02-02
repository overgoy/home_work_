package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func worker(id int, counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
	fmt.Printf("Горутина %d завершила работу\n", id)
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	numWorkers := 10
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, counter, &wg)
	}

	wg.Wait()
	fmt.Printf("Общее количество выполненных задач: %d\n", counter.GetValue())
}
