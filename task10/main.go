package main

import (
	"fmt"
	"math"
)

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
