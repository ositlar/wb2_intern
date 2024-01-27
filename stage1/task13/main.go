package main

import (
	"fmt"
	"reflect"
)

//Поменять местами два числа без создания временной переменной.

func main() {
	a, b := 5, 8
	ReplaceByXOR(a, b)
	Replace(a, b)
	a, b = ReplaceBySwapper([]int{a, b})
	fmt.Println(a, b)
}

func ReplaceByXOR(a, b int) {
	fmt.Printf("XOR:Before: a=%d, b=%d\n", a, b)
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("XOR:After: a=%d, b=%d\n", a, b)
}

func Replace(a, b int) {
	fmt.Printf("Before: a=%d, b=%d\n", a, b)
	a, b = b, a //a, b, a = b, a, b
	fmt.Printf("After: a=%d, b=%d\n", a, b)
}

func ReplaceBySwapper(s []int) (int, int) {
	fmt.Printf("Before swap: %v\n", s)
	swapF := reflect.Swapper(s)
	swapF(0, 1)
	fmt.Printf("After swap: %v\n", s)
	return s[0], s[1]
}
