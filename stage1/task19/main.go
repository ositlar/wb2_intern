package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	str := "начальная строка abcабв123"
	fmt.Println(str)

	result := ReverseString(str)
	fmt.Println(result)
}

func ReverseString(str string) string {
	chars := strings.Split(str, "")
	//bytes := []byte(str)
	n := len(chars)
	for i := 0; i < n/2; i++ {
		chars[i], chars[n-i-1] = chars[n-i-1], chars[i]
	}
	result := strings.Join(chars, "")
	return result
}
