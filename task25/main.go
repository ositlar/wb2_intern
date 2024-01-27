package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.
func main() {
	a := 1
	sleep(500)
	fmt.Println(a)
}

// Принимает время в миллисекундах
func sleep(ms int) {
	end := time.Now().Add(time.Duration(ms) * time.Millisecond)
	for time.Now().Before(end) {
	}
}
