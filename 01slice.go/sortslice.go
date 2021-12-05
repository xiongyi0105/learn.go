package main

import (
	"fmt"
	"sort"
)

// sort包可对切片排序，切片可用切片文字创建，创建时底层先创建数组然后返回切片的引用。
func main() {
	var sliceString = []string{"Python", "Go", "C++", "JAVA"}
	var sliceInt = []int{12, 3, 9, 1, 42, 76, 24, 19}
	sort.Strings(sliceString)
	sort.Ints(sliceInt)
	fmt.Println("NewSliceString:", sliceString)
	fmt.Println("NewSliceInt:", sliceInt)
}
