package main

import "fmt"

func main() {
	var a []int
	// var a = make([]int, 0)
	if a == nil {
		fmt.Println("It's nil", a)
	} else {
		fmt.Println("It's not nil", a)
	}

}
