package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

//Написал два варианта - 1 с пакетом big, 2 через строки

func main() {
	ByPkgBig()
	ByStr()
}

func ByPkgBig() {
	num1 := big.NewInt(1234567890)
	num2 := big.NewInt(198765432)

	sum := big.NewInt(0)
	sum.Add(num1, num2)
	fmt.Println("big - Add: ", sum)

	sub := big.NewInt(0)
	sub.Sub(num1, num2)
	fmt.Println("big - Sub: ", sub)

	multiply := big.NewInt(0)
	multiply.Mul(num1, num2)
	fmt.Println("big - Mul: ", multiply)

	div := big.NewInt(0)
	div.Div(num1, num2)
	fmt.Println("big - Div: ", div)
}

func ByStr() {
	str1 := "1234567890"
	str2 := "198765432"
	sum := StrSum(str1, str2)
	sub := StrSubtract(str1, str2)
	multiply := StrMultiply(str1, str2)
	fmt.Println("str - Sum: ", sum)
	fmt.Println("str - Sub: ", sub)
	fmt.Println("str - Mul: ", multiply)
}

func ToString(arr []string) string {
	reversedArr := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		reversedArr[i] = arr[len(arr)-1-i]
	}

	result := strings.Join(reversedArr, "")
	return result
}

func StrSum(str1, str2 string) string {
	num1 := strings.Split(str1, "")
	num2 := strings.Split(str2, "")
	var minLen int = 0
	if len(num1) > len(num2) {
		minLen = len(num1)
	} else {
		minLen = len(num2)
	}

	result := make([]string, 0)
	var remains int = 0
	var digit1, digit2 int
	for i := 0; i < minLen; i++ {
		if i < len(num1) {
			digit1, _ = strconv.Atoi(num1[len(num1)-i-1])
		} else {
			digit1 = 0
		}
		if i < len(num2) {
			digit2, _ = strconv.Atoi(num2[len(num2)-i-1])
		} else {
			digit2 = 0
		}
		sum := digit1 + digit2 + remains
		remains = sum / 10
		result = append(result, fmt.Sprint(sum%10))
	}
	strResult := ToString(result)
	return strResult
}

func StrSubtract(str1, str2 string) string {
	num1 := strings.Split(str1, "")
	num2 := strings.Split(str2, "")
	var minLen int = 0
	if len(num1) > len(num2) {
		minLen = len(num1)
	} else {
		minLen = len(num2)
	}

	result := make([]string, 0)
	var remains int = 0
	var digit1, digit2 int
	for i := 0; i < minLen; i++ {
		if i < len(num1) {
			digit1, _ = strconv.Atoi(num1[len(num1)-i-1])
		} else {
			digit1 = 0
		}
		if i < len(num2) {
			digit2, _ = strconv.Atoi(num2[len(num2)-i-1])
		} else {
			digit2 = 0
		}
		sub := digit1 - digit2 - remains
		remains = 0
		if sub < 0 {
			sub += 10
			remains = 1
		}
		result = append(result, fmt.Sprint(sub))
	}
	strResult := ToString(result)
	return strResult
}

func StrMultiply(str1, str2 string) string {
	num1 := strings.Split(str1, "")
	num2 := strings.Split(str2, "")
	len1 := len(num1)
	len2 := len(num2)

	result := make([]int, len1+len2)

	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			digit1, _ := strconv.Atoi(num1[i])
			digit2, _ := strconv.Atoi(num2[j])
			result[i+j+1] += digit1 * digit2
		}
	}

	for i := len(result) - 1; i > 0; i-- {
		result[i-1] += result[i] / 10
		result[i] %= 10
	}

	strResult := ""
	for _, digit := range result {
		strResult += fmt.Sprint(digit)
	}

	return strings.TrimLeft(strResult, "0")
}
