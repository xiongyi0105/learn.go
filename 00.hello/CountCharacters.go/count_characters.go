package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	countA := "asSASA ddd dsjkdsjs dk"
	countB := "asSASA ddd dsjkdsjsこん dk"
	countC := "asSASA ddd dsjkdsjsこ嗨 dk"
	fmt.Println(len(countA))
	fmt.Println(utf8.RuneCountInString(countA))
	fmt.Println(len(countB))
	fmt.Println(utf8.RuneCountInString(countB))
	fmt.Println(len(countC))
	fmt.Println(utf8.RuneCountInString(countC))
}
