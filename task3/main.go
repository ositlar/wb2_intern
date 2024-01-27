package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0
	var wg sync.WaitGroup
	wg.Add(len(numbers))
	for i, num := range numbers {
		go func(i, num int) {
			defer wg.Done()
			sum += num * num
		}(i, num)
	}
	wg.Wait()
	fmt.Println("task3: ", sum)
}
