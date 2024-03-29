package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {
	var x interface{} = []int{1, 2, 3}
	var y interface{} = 2.75

	//By fmt
	xType := fmt.Sprintf("%T", x)
	fmt.Println(xType)

	//By reflect
	yType := reflect.TypeOf(y)
	yValue := reflect.ValueOf(y)
	fmt.Println(yType, ": ", yValue)
}
