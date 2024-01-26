package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	fmt.Println("Safe map")
	WithSafeMap()
	time.Sleep(time.Second * 2)
	fmt.Println("\n Sync map")
	WithSyncMap()
}

func WithSafeMap() {
	wg := sync.WaitGroup{}

	sm := &SafeMap{
		concurrentMap: map[int]int{},
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Add(i, i*i)
			log.Print(i, " :add ", i*i)
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value, err := sm.Get(i)
			if err != nil {
				log.Print(err)
			}
			log.Print(i, ": get ", value)
		}(i)
	}

	wg.Wait()
}

func WithSyncMap() {
	var m sync.Map
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i*i)
			log.Print(i, " :add ", i*i)
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value, res := m.Load(i)
			if !res {
				log.Print("Get error in ", i)
			}
			log.Print(i, ": get ", value)
		}(i)
	}

	wg.Wait()
}
