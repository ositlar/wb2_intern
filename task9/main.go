package main

import (
	"fmt"
)

func FirstToSecond() {
	for val := range c1 {
		c2 <- val * 2
	}
}

var (
	c1 = make(chan int)
	c2 = make(chan int)
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	defer close(c1)
	defer close(c2)
	go FirstToSecond()
	for i := range arr {
		c1 <- arr[i]
		fmt.Println(<-c2)
	}
}
