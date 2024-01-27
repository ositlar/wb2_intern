package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var x int64 = 0
	fmt.Println(x)
	for i := 0; i < 64; i++ {
		x = SetBit(x, i)
		fmt.Println(x)
	}

	for i := 0; i < 64; i++ {
		x = SetBit(x, i)
		fmt.Println(x)
	}
}

func SetBit(x int64, i int) int64 {
	x ^= 1 << i
	return x
}
