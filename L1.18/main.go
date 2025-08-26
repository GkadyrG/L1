package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value atomic.Int64
}

func (c *Counter) Increment() {
	c.value.Add(1)
}

func (c *Counter) Value() int64 {
	return c.value.Load()
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	numGoroutines := 1000
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	fmt.Printf("Итоговое значение счетчика: %d\n", counter.Value())

}
