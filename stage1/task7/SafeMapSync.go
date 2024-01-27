package main

import "sync"

type SafeMapSync struct {
	Map sync.Map
}
