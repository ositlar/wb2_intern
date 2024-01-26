package main

import (
	"fmt"
	"strconv"
)

func main() {
	var elems, target int

	fmt.Print("Number of elements? ")
	fmt.Scan(&elems)

	var array = make([]int, elems)

	fmt.Print("Generated array: ")
	for i := 0; i < elems; i++ {
		array[i] = i
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()

	fmt.Println("What are we looking for:")
	fmt.Scan(&target)

	var assumption, result = binarySearch(&array, target)
	if result {
		fmt.Println("Found: " + strconv.Itoa(assumption))
	} else {
		fmt.Println("Nothing found!")
	}
}

func binarySearch(array *[]int, target int) (int, bool) {
	fmt.Println("Looking for...")
	var mid, assumption int

	min := 0
	high := len(*array) - 1

	for min <= high {
		mid = (min + high) / 2
		assumption = mid
		if assumption == target {
			return assumption, true
		}
		if assumption > target {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}
	return 0, false
}
