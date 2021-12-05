package main

import "fmt"

func main() {
	var onedimendslice = make([]int, 10)
	var multislice = make([][]int, 10)
	for i := range multislice {
		multislice[i] = make([]int, 10)
	}

	fmt.Println(onedimendslice)
	fmt.Println(multislice)
}
