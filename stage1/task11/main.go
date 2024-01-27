package main

import (
	"fmt"
	"sort"
	"time"
)

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	var set1 = []int{1, 45, 6, 2, 8, 11, 324, 5556, 12, 123, 546}
	var set2 = []int{4, 42, 6, 2, 434, 67, 1, 4567, 33, 123, 546}
	var result = []int{}

	start := time.Now()
	result = smallSizeSet(set1, set2)
	end := time.Now()
	fmt.Println(result, "time: ", end.Sub(start))

	start = time.Now()
	result = BigSizeSet(set1, set2)
	end = time.Now()
	fmt.Println(result, "time: ", end.Sub(start))
}

func smallSizeSet(set1, set2 []int) []int {
	var result = []int{}
	for i := 0; i < len(set1); i++ {
		for j := 0; j < len(set2); j++ {
			if set1[i] == set2[j] {
				result = append(result, set1[i])
			}
		}
	}

	return result
}

func BigSizeSet(set1, set2 []int) []int {
	var result = []int{}

	sort.Slice(set1, func(i, j int) bool { return set1[i] < set1[j] })
	sort.Slice(set2, func(i, j int) bool { return set2[i] < set2[j] })

	var minLen int
	if len(set1) > len(set2) {
		minLen = len(set2)
	} else {
		minLen = len(set1)
	}

	for i := 0; i < minLen; i++ {
		target := sort.SearchInts(set2, set1[i])
		result = append(result, target)
	}

	return result
}
