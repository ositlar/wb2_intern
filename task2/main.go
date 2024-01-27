package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	var1()
	var2()
}

func var1() {
	numbers := []int{2, 4, 6, 8, 10}
	squares := make([]int, len(numbers))

	for i, num := range numbers {
		go func(i, num int) {
			squares[i] = num * num
		}(i, num)
	}

	fmt.Println("var1: ", squares)
}

/////////////Второй вариант решения///////////////

func var2() {
	numbers := []int{2, 4, 6, 8, 10}
	squares := make([]int, len(numbers))

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for i, num := range numbers {
		go func(i, num int) {
			defer wg.Done()
			squares[i] = num * num
		}(i, num)
	}

	wg.Wait()

	fmt.Println("var2: ", squares)
}
