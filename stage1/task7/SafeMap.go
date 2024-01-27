package main

import (
	"errors"
	"sync"
)

type SafeMap struct {
	mx            sync.RWMutex
	concurrentMap map[int]int
}

func (sm *SafeMap) Add(key, digit int) {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	sm.concurrentMap[key] = digit
}

func (sm *SafeMap) Get(key int) (int, error) {
	sm.mx.RLock()
	defer sm.mx.RUnlock()

	if value, ok := sm.concurrentMap[key]; ok {
		return value, nil
	}

	return 0, errors.New("Digit doesnt exists")
}
