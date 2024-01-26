package main

import "fmt"

func main() {
	strArr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Print(ArrToSet(strArr))
}

func ArrToSet(strArr []string) []string {
	Map := make(map[string]bool)
	set := make([]string, 0, len(Map))
	for _, s := range strArr {
		Map[s] = true
	}
	for k := range Map {
		set = append(set, k)
	}
	return set
}
