package main

import "fmt"

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка
//Либо можно использовать sort (11 задание, функция BigSizeSearch)

func main() {
	arrToSort := []int{5, 8, 2, 6, 89, 2, 35, 5678, 21, 45, 234, 78, 12, 53, 7869, 24, 56, 123, 4, 6789, 345}
	fmt.Println(arrToSort)
	result := quickSortStart(arrToSort)
	fmt.Println(result)
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}
