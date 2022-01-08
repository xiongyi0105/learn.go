package main

import (
	"fmt"
)

type Simpler interface {
}

var RSimple = [1]int{13}

type TSimple struct {
	hello string
}

func fi(items ...interface{}) {
	for _, x := range items {
		switch x.(type) {
		case TSimple:
			fmt.Println("TSimple")
		case [1]int:
			fmt.Println("arry")
		default:
			fmt.Println("no")
		}
	}
}

func main() {
	var Simp1 Simpler
	var Simp2 Simpler
	Simp1 = RSimple
	Simp2 = TSimple{"123"}
	fi(Simp1, Simp2)
}
