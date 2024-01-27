package main

import "fmt"

//Разработать программу, которая проверяет, что все символы в строке уникальные
//(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

func main() {
	fmt.Println("abcd - ", CheckUnique("abcd"))
	fmt.Println("abCdefAaf - ", CheckUnique("abCdefAaf"))
	fmt.Println("aabcd - ", CheckUnique("aabcd"))
}

func CheckUnique(str string) bool {
	Map := make(map[rune]bool)
	for _, char := range str {
		_, ok := Map[char]
		if ok {
			return false
		}
		Map[char] = true
	}
	return true
}
