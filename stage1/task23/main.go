package main

import "fmt"

//Удалить i-ый элемент из слайса.

func main() {
	slice := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Println(slice)
	var i int
	fmt.Println("Введите индекс элемента для удаления")
	fmt.Scan(&i)
	FastNotOrdered(i)
	SlowOrdered(i)

}

func FastNotOrdered(i int) {
	slice := []string{"A", "B", "C", "D", "E", "F"}
	slice[i] = slice[len(slice)-1]
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}

func SlowOrdered(i int) {
	slice := []string{"A", "B", "C", "D", "E", "F"}
	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
