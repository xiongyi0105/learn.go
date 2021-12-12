package main

import "fmt"

func main() {
	myslice := make([]int, 5, 10)
	fmt.Println(myslice)
	myslice = myslice[0:9]
	for i := 0; i < len(myslice); i++ {
		myslice[i] = i
	}

	//myslice = myslice[0:9]
	fmt.Println(myslice)
	sl := myslice[3:3]
	fmt.Println(len(sl), sl)
	sl2 := myslice[8:9]
	fmt.Println(len(sl2), sl2)
}
