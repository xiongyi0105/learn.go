package main

import "fmt"

func main() {
	initList := []int{1, 4, 7, 3, 9, 2, 8}
	for i, j := 0, len(initList)-1; i < j; i, j = i+1, j-1 {
		initList[i], initList[j] = initList[j], initList[i]

	}
	fmt.Println(initList)
}
