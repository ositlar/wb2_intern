package main

import (
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	sm := &SafeMap{
		concurrentMap: map[int]int{},
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Add(i, i*i)
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value, err := sm.Get(i)
			if err != nil {
				log.Fatal(err)
			}
			log.Print(value)
		}(i)
	}

	wg.Wait()
}
