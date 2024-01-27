package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	str := "snow dog sun now order changed"
	fmt.Println(str)
	result := ReverseWordsOrder(str)
	fmt.Println(result)

}

func ReverseWordsOrder(str string) string {
	words := strings.Split(str, " ")
	result := ""
	for i := range words {
		result += words[len(words)-i-1] + " "
	}
	return result
}
