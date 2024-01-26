package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Можно и без mutex (будет быстрее),
// но если во время выполнения нужно считывать этот счетчик, значение может отличаться от того, которое должно быть
type ConcurrentCounter struct {
	m       sync.Mutex
	counter int64
}

func (c *ConcurrentCounter) Inc() {
	atomic.AddInt64(&c.counter, 1)
}

func main() {
	counter := &ConcurrentCounter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			counter.m.Lock()
			defer counter.m.Unlock()
			counter.Inc()
			fmt.Println(counter.counter)
			wg.Done()
		}()
	}
	wg.Wait()
}
