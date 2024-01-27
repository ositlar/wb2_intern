package main

import (
	"fmt"
	"math"
)

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	temperature := []float32{-25.4, -27, 13, 19, 15.5, 24.5, -21, 32.5}
	set := make(map[int][]float32)
	var min, max float32 = temperature[0], temperature[0]
	for i := range temperature {
		if temperature[i] < min {
			min = temperature[i]
		}
		if temperature[i] > max {
			max = temperature[i]
		}
	}
	var j = int(math.Round(float64(min)/10) * 10)

	for ; j <= int(math.Round(float64(max)/10)*10); j += 10 {
		tempArr := make([]float32, 0)
		for i := range temperature {
			if temperature[i] >= float32(j) && temperature[i] < float32(j+10) {
				tempArr = append(tempArr, temperature[i])
			}
		}
		if len(tempArr) == 0 {
			continue
		}
		set[j] = append(set[j], tempArr...)
	}

	fmt.Println(set)
}
