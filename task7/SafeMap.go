package main

import (
	"errors"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	concurrentMap map[int]int
}

func (sm *SafeMap) Add(key, digit int) {
	sm.Lock()
	defer sm.Unlock()
	sm.concurrentMap[key] = digit
}

func (sm *SafeMap) Get(key int) (int, error) {
	sm.RLock()
	defer sm.RUnlock()

	if value, ok := sm.concurrentMap[key]; ok {
		return value, nil
	}

	return 0, errors.New("Digit doesnt exists")
}
